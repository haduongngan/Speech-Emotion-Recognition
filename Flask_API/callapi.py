import requests
from google.cloud import speech_v1p1beta1 as speech

import os
os.environ["GOOGLE_APPLICATION_CREDENTIALS"]="/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/analog-crossing-353415-448731397826.json"

def token_wav(path):

    client = speech.SpeechClient()

    speech_file = path

    with open(speech_file, "rb") as audio_file:
        content = audio_file.read()

    audio = speech.RecognitionAudio(content=content)

    diarization_config = speech.SpeakerDiarizationConfig(
        enable_speaker_diarization=True,
        min_speaker_count=2,
        max_speaker_count=10,
    )

    config = speech.RecognitionConfig(
        encoding=speech.RecognitionConfig.AudioEncoding.LINEAR16,
        sample_rate_hertz=8000,
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
    for word_info in words_info:
        print(
            u"word: '{}', speaker_tag: {}".format(word_info.word, word_info.speaker_tag)
        )



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






# token_wav('/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/6.wav')
# speech_to_text('/Users/nguyentrongdat/Documents/19021240_NguyenTrongDat/6.wav')

