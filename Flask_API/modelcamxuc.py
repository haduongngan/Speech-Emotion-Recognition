import torch.nn as nn

import pickle
import librosa
import numpy as np
import torch
class CNN(nn.Module):

    def __init__(self, ):
        super(CNN, self).__init__()

        # Block #1:
        self.layer1 = nn.Sequential(
            nn.Conv1d(in_channels=1, out_channels=256, kernel_size=5),
            nn.ReLU(),
            nn.MaxPool1d(kernel_size=5, stride=2)
        )

        # Block #2:
        self.layer2 = nn.Sequential(
            nn.Conv1d(in_channels=256, out_channels=256, kernel_size=5),
            nn.ReLU(),
            nn.MaxPool1d(kernel_size=5, stride=2)
        )

        # Block #3:
        self.layer3 = nn.Sequential(
            nn.Conv1d(in_channels=256, out_channels=128, kernel_size=5),
            nn.ReLU(),
            nn.MaxPool1d(kernel_size=5, stride=2)
        )

        # Block #4:
        self.layer4 = nn.Sequential(
            nn.Conv1d(in_channels=128, out_channels=64, kernel_size=5),
            nn.ReLU(),
            nn.MaxPool1d(kernel_size=5, stride=2)
        )

        # Block #5:
        self.layer5 = nn.Sequential(
            nn.Linear(in_features=192, out_features=32),
            nn.ReLU(),
            nn.Dropout(p=0.3)
        )

        # FC 5 → softmax
        self.fc = nn.Linear(in_features=32, out_features=8)
        self.softmax = nn.Softmax(dim=1)

    def forward(self, x):
        # Channel x H = 1 x 162
        out = self.layer1(x.view(-1, 1, 162))

        out = self.layer2(out)
        out = self.layer3(out)
        out = self.layer4(out)

        out = out.view(out.size(0), -1)
        out = self.layer5(out)
        out = self.fc(out)
        out = self.softmax(out)
        return out

def extract_features(data, sample_rate):
    # ZCR
    result = np.array([])
    zcr = np.mean(librosa.feature.zero_crossing_rate(y=data).T, axis=0)
    result = np.hstack((result, zcr))

    # MFCC
    mfcc = np.mean(librosa.feature.mfcc(y=data, sr=sample_rate).T, axis=0)
    result = np.hstack((result, mfcc))

    # Log Mel-Spectrogram
    mel = np.mean(librosa.feature.melspectrogram(y=data, sr=sample_rate).T, axis=0)
    result = np.hstack((result, mel))

    # Chroma
    chroma_stft = np.mean(librosa.feature.chroma_stft(S=np.abs(librosa.stft(data)), sr=sample_rate).T, axis=0)
    result = np.hstack((result, chroma_stft))

    # Root Mean Square Value
    rms = np.mean(librosa.feature.rms(y=data).T, axis=0)
    result = np.hstack((result, rms))

    return result


Emo = ['neutral', 'calm', 'happy', 'sad', 'angry', 'fear', 'disgust', 'surprise']


def emotion_recognition(audio_file):
    trained_model = pickle.load(open('models/model_162.pkl', 'rb'))
    scaler = pickle.load(open('models/scaler.pkl', 'rb'))

    # load audio files with librosa
    data, sample_rate = librosa.load(audio_file)
    feat = extract_features(data, sample_rate)
    feat = np.array(feat)
    feat = feat[None, :]
    sc_feat = scaler.transform(feat)
    sc_feat = torch.from_numpy(sc_feat.astype('float32'))
    sc_feat = sc_feat.cpu()
    trained_model = trained_model.cpu()
    prediction = trained_model(sc_feat)
    pred = torch.argmax(prediction, dim=1)
    return Emo[pred]


