from llama_index.llms.openai import OpenAI
from config import app_settings
from Base import Dialog, RouteAnswer
from Prompts import PromptClass

llm = OpenAI(model=app_settings.MODEL, api_key=app_settings.API_KEY, api_base=app_settings.API_BASE)

def get_dialogue(text, n):
    restaurant_obj = (
        llm.as_structured_llm(Dialog)
        .complete(PromptClass.prompt_tmpl_link.format(link=text, n=n))
        .raw
    )
    return restaurant_obj

def get_router_result(text):
    # print(text)
    prompt = PromptClass.ROUTER_PROMPT.format(query_str=text)
    print(prompt)
    router_obj = (
        llm.as_structured_llm(RouteAnswer)
        .complete(prompt)
        .raw
    )
    print(router_obj.result)
    return router_obj.result