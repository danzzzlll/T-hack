import sys
sys.path.append("voicekit_examples/python")


from voicekit_examples.python.tinkoff.cloud.tts.v1 import tts_pb2_grpc, tts_pb2
import grpc
import os
import wave
from tqdm import tqdm
import glob

from config import app_settings

class Text2Speech:

    def __init__(self, dialogue):

        self.dialogue = dialogue
        self.voices = {
            'father': app_settings.FATHER_NAME,
            'daughter': app_settings.D_NAME
        }
        self.pitch = {
            'father': app_settings.PITCH_FATHER,
            'daughter': app_settings.PITCH_D
        }
        self.sample_rate = 16000
        self.token = app_settings.SPEECH_KEY
        self.metadata = [("authorization", f"Bearer {self.token}")]
        self.endpoint = app_settings.SPEECH_ENDPOINT
        self.audio_path = 'audios/'
    
    def get_part(self):
        
        for idx, part in tqdm(enumerate(self.dialogue)):
            request = self.build_request(
                self.voices[part['role']], 
                self.pitch[part['role']], 
                part['phrase']
            )
            self.save(request, f"{self.audio_path}/part_{idx}_{part['role']}.wav")
    
    def build_request(self, role, pitch, text):
         
        stub = tts_pb2_grpc.TextToSpeechStub(grpc.secure_channel(self.endpoint, grpc.ssl_channel_credentials()))
        request = tts_pb2.SynthesizeSpeechRequest(
            input=tts_pb2.SynthesisInput(
                text=text,
            ),
            audio_config=tts_pb2.AudioConfig(
                audio_encoding=tts_pb2.LINEAR16,
                sample_rate_hertz=self.sample_rate,
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

