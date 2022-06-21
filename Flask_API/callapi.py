from unittest import result
from collections import Counter
import librosa
import requests
from google.cloud import speech_v1p1beta1 as speech
from pydub import AudioSegment
from sox import Transformer
import soundfile
import moviepy.editor as moviepy
import pandas as pd
import numpy as np
from pydub import AudioSegment
import os
from modelcamxuc import *
# os.environ["GOOGLE_APPLICATION_CREDENTIALS"]="/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/analog-crossing-353415-448731397826.json"
os.environ["GOOGLE_APPLICATION_CREDENTIALS"]="speech.json"
SAMPLE_RATE_HERTZ = 22050

# convert ogg file to wav file with 16 bit
def convert_file(file1, file2):
    data, samplerate = soundfile.read(file1)
    soundfile.write(file2, data, samplerate, subtype='PCM_16')
# convert_file("1.ogg","new.wav")

def token_wav(path):
    # read sample_rate and assign to config
    _, samplerate = soundfile.read(path)

    client = speech.SpeechClient()

    speech_file = path

    with open(speech_file, "rb") as audio_file:
        content = audio_file.read()

    audio = speech.RecognitionAudio(content=content)

    diarization_config = speech.SpeakerDiarizationConfig(
        enable_speaker_diarization=True,
        min_speaker_count=2,
        max_speaker_count=2,
    )

    config = speech.RecognitionConfig(
        encoding=speech.RecognitionConfig.AudioEncoding.LINEAR16,
        sample_rate_hertz=samplerate,
        language_code="en-US",
        diarization_config=diarization_config,
    )

    print("Waiting for operation to complete...")
    response = client.recognize(config=config, audio=audio)

    # The transcript within each result is separate and sequential per result.
    # However, the words list within an alternative includes all the words
    # from all the results thus far. Thus, to get all the words with speaker
    # tags, you only have to take the words list from the last result:
    result = response.results[-1]

    words_info = result.alternatives[0].words

    # Printing out the output:
    list_word = []
    for word_info in words_info:
        seconds_start = word_info.start_time.total_seconds()
        seconds_end = word_info.end_time.total_seconds()
        duration = round(seconds_end - seconds_start,2)
        print(
            u"word: '{}', speaker_tag: {}, start_time: {}, end_time: {}, duration : {}".format(word_info.word, word_info.speaker_tag, seconds_start, seconds_end, duration)
        )
        list_word.append([word_info.word, word_info.speaker_tag, seconds_start, seconds_end, duration])
    return list_word





def speech_to_text(AUDIO_PATH):
    url = "https://viettelgroup.ai/voice/api/asr/v1/rest/decode_file"
    headers = {
        'token': 'anonymous',
        # 'sample_rate': 16000,
        # 'format':'S16LE',
        # 'num_of_channels':1,
        # 'asr_model': 'model code'
    }

    # CERT_PATH ='/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/wwwvtccai.crt'
    s = requests.Session()
    files = {'file': open(AUDIO_PATH, 'rb')}
    response = requests.post(url, files=files, headers=headers)

    print(response.text)

# segment
# input : dataframe['words','speaker_tag','start_time','end_time','duration']
# output : dataframe['start_time','end_time','speaker']
GAP_MAX = 0.1
TOTAL_MAX = 3


def segment(df):
    segment = []
    speaker_staff = df['speaker_tag'][0]

    id_segment = 1
    start_time = df['start_time'][0]
    end_time = df['end_time'][0]
    id_speaker = df['speaker_tag'][0]
    total_time = df['duration'][0]

    result = np.array([])
    result = np.hstack((result, start_time))

    for i in range(df.shape[0]):
        if (i == df.shape[0] - 1):
            result = np.hstack((result, end_time))
            id = int(id_speaker)
            result = np.hstack((result, id))
            segment.append(result)

        elif (id_speaker != df['speaker_tag'][i + 1] or (
                df['start_time'][i + 1] - end_time) > GAP_MAX or total_time > TOTAL_MAX):
            result = np.hstack((result, end_time))
            id = int(id_speaker)
            result = np.hstack((result, id))
            segment.append(result)

            id_segment = id_segment + 1
            result = np.array([])
            start_time = df['start_time'][i + 1]
            result = np.hstack((result, start_time))
            end_time = df['end_time'][i + 1]
            total_time = df['duration'][i + 1]
            id_speaker = df['speaker_tag'][i + 1]

        else:
            end_time = df['end_time'][i + 1]
            total_time = total_time + df['duration'][i + 1]

    segment_df = pd.DataFrame(segment, columns=['start_time', 'end_time', 'speaker'])
    return segment_df, speaker_staff

# token_wav('/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/6.wav')
# speech_to_text('/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/6.wav')




def get_token_wav(path):
    # test call api and add to dataframe
    words = token_wav(path)
    df = pd.DataFrame(words, columns=['words', 'speaker_tag', 'start_time', 'end_time', 'duration'])
    df.to_csv('df.csv', sep="\t", encoding='utf-8')
    # print(df)

    segment_df, speaker_staff = segment(df)
    print(segment_df)
    print(speaker_staff)
    return segment_df, speaker_staff

def get_dur_audio(path):
    dur = librosa.get_duration(filename=path)
    return dur


def split_wav_file(path, t1, t2, id_speaker):

    t1 = t1 * 1000  # Works in milliseconds
    t2 = t2 * 1000
    newAudio = AudioSegment.from_wav(path)
    newAudio = newAudio[t1:t2]
    ext_path = 'audio_speaker/' + str(id_speaker) + '/' + str(t1) + 'newSong.wav'
    newAudio.export(ext_path, format="wav")



def split_wav_file_df(path,df):
    for index, row in df.iterrows():
        t1 = float(row[0])
        t2 = float(row[1])
        id = int(row[2])
        split_wav_file(path,t1,t2,id)



def get_split(path):
    segment_df, staff = get_token_wav(path)
    split_wav_file_df(path,segment_df)
    return staff


from collections import Counter
def get_all_emotion_recognition_one_people(path):
    emotion = []
    count = 0
    for i in Emo:
        emotion.append(i)
    for i in os.listdir(os.path.join('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker',path)):
        audio_path = os.path.join('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker',path,i)
        emo = emotion_recognition(audio_path)
        emotion.append(emo)
        count+=1

    emotion = Counter(emotion)
    # for i in emotion:
    #     emotion[i] = int(emotion[i]) - 1
    for i in emotion:
        emotion[i] = round(emotion[i]/count,2)
    return emotion


# token_wav('1.wav')