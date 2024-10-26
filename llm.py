from llama_index.llms.openai import OpenAI
from config import app_settings
from Base import Dialog
from Prompts import PromptClass

llm = OpenAI(model=app_settings.MODEL, api_key=app_settings.API_KEY, api_base=app_settings.API_BASE)

restaurant_obj = (
    llm.as_structured_llm(Dialog)
    .complete(PromptClass.prompt_tmpl_link.format(link="https://software.tbank.ru/docs/voicekit/ttstutorial", n=app_settings.YEARS))
    .raw
)