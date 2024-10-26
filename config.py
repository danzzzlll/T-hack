import os
from dotenv import load_dotenv
from pydantic import Field
from pydantic_settings import BaseSettings

class Config(BaseSettings):
    API_KEY: str = Field('sk-vfaXetQTUOUoffJ0H9oXhsRBdtfLRI5n')
    MODEL: str = Field('gpt-4o-mini')
    API_BASE: str = Field('https://api.proxyapi.ru/openai/v1')
    YEARS: int = Field(16)
    LINK: str = Field("https://software.tbank.ru/docs/voicekit/ttstutorial")
    # Speech
    SPEECH_ENDPOINT: str = Field("api.tinkoff.ai:443")
    SPEECH_KEY: str = Field("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InIxUGNSUGsvMVo3WG9QSGxIS3d2cmdWUkxnQ1ZFTnByRHZPK1ArODM2NHM9VFRTX1RFQU0ifQ.eyJpc3MiOiJ0ZXN0X2lzc3VlciIsInN1YiI6InRlc3RfdXNlciIsImF1ZCI6InRpbmtvZmYuY2xvdWQudHRzIiwiZXhwIjoxNzI5OTkxNDA2LjB9.fC9xzEoZ6izg77Beph8Kbjz7A-VHvImgsvhD8y--aDY")
    FATHER_NAME: str = Field("dorofeev")
    D_NAME: str = Field("vika")

    PITCH_FATHER: float = Field(1.)
    PITCH_D: float = Field(1.7)


    class Config:
        env_file = "env"  


app_settings = Config()
print(app_settings)

