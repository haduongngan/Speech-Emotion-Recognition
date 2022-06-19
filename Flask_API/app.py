import json


import urllib.request
from model import *
import audioread
from flask import Flask, request #import main Flask class and request object
from flask_cors import CORS
import logging

logging.basicConfig(level=logging.INFO)

logger = logging.getLogger('HELLO WORLD')

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model = load_model('models/resnet_50_0605_1007_r2.t7')
model = model.to(device)
app = Flask(__name__)
cors = CORS(app)

# load model


# @app.route('/')
# def landing():
#     return 'Landing page'
#


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
    inputs = inputs[None,:]
    inputs = inputs.to(device)
    outputs = model(inputs)
    outputs = list(outputs.cpu())
    prediction_data = max(outputs)
    prediction_data=outputs.index(prediction_data)
    return prediction_data




def getjsonreturn():
    gender = prediction_gender('1.wav')
    if(gender ==1):
        gender = 'male'
    else:
        gender = 'female'
    return gender




# Get the blob of type "audio/webm;codecs=opus"
@app.route('/audio_record', methods=['GET', 'POST'])
def save_record():
    print('Done')
    file = request.files['file']
    try:
        audio = request.files["file"]
        path = '1.wav'
        audio.save(path)
    except:
        print('B')






@app.route('/')
def hello_world():
    return 'Hello, World!'





# if __name__ == "__main__":
#     app.run(debug=True)
