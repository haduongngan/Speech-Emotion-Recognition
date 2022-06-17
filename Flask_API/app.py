import torch
from flask import Flask, request
from flask_cors import CORS, cross_origin
from model import *
import audioread
from flask import Flask, request #import main Flask class and request object
from flask_cors import CORS
import logging

logging.basicConfig(level=logging.INFO)

logger = logging.getLogger('HELLO WORLD')

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
app = Flask(__name__)
cors = CORS(app)

# load model
model = load_model('/Users/nguyentrongdat/Desktop/Speechprocessing/Flask_API/models/resnet_50_0605_1007_r2.t7')

model = model.to(device)

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
    inputs = torch.Tensor(feats)
    inputs = inputs[None,:]
    inputs = inputs.to(device)
    outputs = model(inputs)
    outputs = list(outputs.cpu())
    prediction_data = max(outputs)
    prediction_data=outputs.index(prediction_data)
    return prediction_data






# Get the blob of type "audio/webm;codecs=opus"
@app.route('/audio_record', methods=['POST'])
def save_record():
    logger.info("welcome to upload`")
    # file = request.files['file']
    #filename = secure_filename(file.title)

    file = request.form['file']

    print('File from the POST request is: {}'.format(file))
    try:
        read_audio_file(file[0])
        return "****** Audio Read ******"
    except:
        print("In the except", file[0]) # Gets printed as undefined
        title = request.form['title']
        print(title) # Able to print title
        return "Request received and responded"
        # app.logger.debug(request.files['file'].filename)





# @app.route("/predict", methods=["GET", "POST"])
# def predictcrop():
#     try:
#         if request.method == "POST":
#             form_values = request.form.to_dict()
#             column_names = ["N", "P", "K", "temperature", "humidity", "ph", "rainfall"]
#             input_data = np.asarray([float(form_values[i].strip()) for i in column_names]).reshape(
#                 1, -1
#             )
#             prediction_data = crop_prediction(input_data)
#             json_obj = json.dumps(prediction_data, default=convert)
#             return json_obj
#     except:
#         return json.dumps({"error":"Please Enter Valid Data"}, default=convert)





@app.route('/')
def hello_world():
    return 'Hello, World!'





if __name__ == "__main__":
    app.run(debug=True)
