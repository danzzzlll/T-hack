import os
from dotenv import load_dotenv
from pydantic import Field
from pydantic_settings import BaseSettings


class Config(BaseSettings):
    API_KEY: str = Field("sk-vfaXetQTUOUoffJ0H9oXhsRBdtfLRI5n")
    MODEL: str = Field('gpt-4o-mini')
    API_BASE: str = Field('https://api.proxyapi.ru/openai/v1')
    YEARS: int = Field(16)
    LINK: str = Field("")

    class Config:
        env_file = ".env"  


app_settings = Config()

