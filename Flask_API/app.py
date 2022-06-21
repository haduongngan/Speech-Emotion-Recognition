import json

import urllib.request
from model import *
import audioread
from flask import Flask, request  # import main Flask class and request object
from flask_cors import CORS
import logging
from callapi import *

logging.basicConfig(level=logging.INFO)

logger = logging.getLogger('HELLO WORLD')

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model = load_model('models/resnet_50_0605_1007_r2.t7')
model = model.to(device)
app = Flask(__name__)
cors = CORS(app)


def read_audio_file(audio_from_post):
    print("Tring to read audio..")
    with audioread.audio_open(audio_from_post) as f:
        print(f.channels, f.samplerate, f.duration)
        for buf in f:
            print(buf)


def prediction_gender(wav):
    model.eval()
    feats = gen_spec(wav, duration=2)
    inputs = torch.tensor(feats)
    inputs = inputs[None, :]
    inputs = inputs.to(device)
    outputs = model(inputs)
    outputs = list(outputs.cpu())
    prediction_data = max(outputs)
    prediction_data = outputs.index(prediction_data)
    if prediction_data == 1:
        prediction_data = 'male'
    else:
        prediction_data = 'female'
    return prediction_data


def delete_all_file():
    import shutil
    shutil.rmtree('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/1')
    shutil.rmtree('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/2')
    shutil.rmtree('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/3')
    os.mkdir('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/1')
    os.mkdir('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/2')
    os.mkdir('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker/3')


def getjsonreturn():
    gender = prediction_gender('1.wav')
    if (gender == 1):
        gender = 'male'
    else:
        gender = 'female'
    return gender


def get_gerder_one_people(path):
    for i in os.listdir(os.path.join('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker', path)):
        audio_path = os.path.join('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/audio_speaker', path, i)
        return prediction_gender(audio_path)
        break


@app.route('/audio_record', methods=['GET', 'POST'])
def save_record():
    print('Done')
    file = request.files['file']
    try:
        audio = request.files["file"]
        path = '1.wav'
        audio.save(path)
        staff = get_split(path)
        dur = get_dur_audio(path)
        dur = round(dur, 2)
        dur = str(dur)
        gender1 = get_gerder_one_people('1')
        gender2 = get_gerder_one_people('2')
        emo1 = get_all_emotion_recognition_one_people('1')
        emo2 = get_all_emotion_recognition_one_people('2')
        feel1 = max(emo1)
        feel2 = max(emo2)

        if staff == 1:
            nhanvien = {
                'gender': gender1,
                'emo': emo1,
                'feel': feel1
            }
            khachhang = {
                'gender': gender2,
                'emo': emo2,
                'feel': feel2
            }
        else:
            nhanvien = {
                'gender': gender2,
                'emo': emo2,
                'feel': feel2
            }
            khachhang = {
                'gender': gender1,
                'emo': emo1,
                'feel': feel1
            }

        data = {
            'customer': khachhang,
            'staff': nhanvien,
            'dur': dur
        }

        print(data)
        json_obj = json.dumps(data)
        print(json_obj)
        # delete_all_file()
        return json_obj
    except:
        return "Error"

    return 'Done'


@app.route('/firstcallanalystic', methods=['GET', 'POST'])
def save():
    print('Done')
    file = request.files['file']
    print(file)
    try:
        audio = request.files["file"]
        path = '2.wav'
        audio.save(path)
        dur = get_dur_audio(path)
        print(dur)
        dur = round(dur, 2)
        dur = str(dur)
        gender = prediction_gender(path)
        print(gender)
        # splitwav(path, dur)
        emo = get_all_emotion_recognition_one_people('1')
        # emo = 'null'
        feel = max(emo)
        # feel = 'null'
        nhanvien = {
            'gender': gender,
            'emo': emo,
            'feel': feel

        }



        data = {
            'customer': nhanvien,
            'dur': dur,
        }
        print(data)
        json_obj = json.dumps(data)
        print(json_obj)
        # delete_all_file()
        return json_obj

    except:
        return "Error"

    return 'Done'


def splitwav(path, dur):
    for i in range(1, dur - 3, 3):
        print(i)
        t1 = i
        t2 = i + 3
        id = 3
        split_wav_file(path, t1, t2, id)



@app.route('/')
def hello_world():
    return 'Hello, World!'


if __name__ == "__main__":
    app.run(debug=True)
