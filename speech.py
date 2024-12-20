import sys
sys.path.append("voicekit_examples/python")


from voicekit_examples.python.tinkoff.cloud.tts.v1 import tts_pb2_grpc, tts_pb2
import grpc
import os
import wave
from tqdm import tqdm
import glob
from pydub import AudioSegment
import random

from config import app_settings

class Text2Speech:

    def __init__(self, dialogue, age):
        
        self.dialogue = dialogue
        self.params = {
            'father': {
                'pitch': 1,
                'voice': 'dorofeev',
                'speed': 1
            },
            'daughter': {
                10: {
                    'pitch': 1.4,
                    'voice': 'vika',
                    'speed': 0.9
                },
                14:{
                    'pitch': 1.3,
                    'voice': 'sveta',
                    'speed': 0.9
                },
                18:{
                    'pitch': 1.2,
                    'voice': 'nika',
                    'speed': 1
                }
            }
        }
        self.age = age
        self.sample_rate = 16000
        self.token = app_settings.SPEECH_KEY
        self.metadata = [("authorization", f"Bearer {self.token}")]
        self.endpoint = app_settings.SPEECH_ENDPOINT
        self.audio_path = 'audios/'

        self.clean_folder()
    
    def get_part(self):

        for idx, part in enumerate(self.dialogue):

            if part['role'] == 'father':
                pitch = self.params[part['role']]['pitch']
                voice = self.params[part['role']]['voice']
                speed = self.params[part['role']]['speed']
            else:
                pitch = self.params[part['role']][self.age]['pitch']
                voice = self.params[part['role']][self.age]['voice']
                speed = self.params[part['role']][self.age]['speed']
            print(pitch, voice, speed)
            request = self.build_request(
                voice, 
                pitch, 
                speed,
                part['phrase']
            )
            self.save(request, f"{self.audio_path}/part_{idx}_{part['role']}.wav")

        return
    
    def build_request(self, role, pitch, speed, text):

        stub = tts_pb2_grpc.TextToSpeechStub(grpc.secure_channel(self.endpoint, grpc.ssl_channel_credentials()))
        request = tts_pb2.SynthesizeSpeechRequest(
            input=tts_pb2.SynthesisInput(
                ssml=text
            ),
            audio_config=tts_pb2.AudioConfig(
                audio_encoding=tts_pb2.LINEAR16,
                sample_rate_hertz=self.sample_rate,
                speaking_rate = speed,
                pitch = pitch
            ),
            voice=tts_pb2.VoiceSelectionParams(
                name=role,
            ),
        )
        response = stub.Synthesize(request, metadata=self.metadata)
        
        return response

    def save(self, response, path):
        with wave.open(path, "wb") as f:
            f.setframerate(self.sample_rate)
            f.setnchannels(1)
            f.setsampwidth(2)
            f.writeframes(response.audio_content)

        return

    def combine(self):
        
        files = list(filter(os.path.isfile, glob.glob(self.audio_path + "*.wav")))
        files.sort(key=lambda x: os.path.getmtime(x))
        outfile = f"{self.audio_path}/dialogue.wav"

        data = []
        for file in files:
            w = wave.open(file, 'rb')
            data.append( [w.getparams(), w.readframes(w.getnframes())] )
            w.close()
            
        output = wave.open(outfile, 'wb')
        output.setparams(data[0][0])
        for i in range(len(data)):
            output.writeframes(data[i][1])
        output.close()

        return
    
    def post_processing(self, dialogue, noise_level=16):
        print('start post processing')
        dialogue = AudioSegment.from_file(dialogue)
        background_noise1 = AudioSegment.from_file("data/background1.wav")
        background_noise2 = AudioSegment.from_file("data/background2.wav")
        background_noise3 = AudioSegment.from_file("data/background3.wav")
        background_noises = [background_noise1, background_noise2, background_noise3]
        background_noises = [noise - noise_level for noise in background_noises]  
        background_noise = random.choice(background_noises)
        background_noise = background_noise[:len(dialogue)]
        final_audio = dialogue.overlay(background_noise)
        final_audio.export("audios/final_output.wav", format="wav")

    def clean_folder(self):
        
        files = os.listdir(self.audio_path)
        for file in files:
            file_path = os.path.join(self.audio_path, file)
            if os.path.isfile(file_path):
                os.remove(file_path)

        return
