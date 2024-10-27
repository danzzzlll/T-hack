import os
import streamlit as st
from openai import OpenAI
from utils import load_and_extract_text
from llm import get_dialogue, get_router_result
from Prompts import PromptClass
from speech import Text2Speech
from utils import get_md_style, display_message


get_md_style()


def main():
    os.makedirs("audios", exist_ok=True)
    st.title("Диалоги о Главном: Отец и Дочь")
    option_age = st.selectbox(
        "Выберете возвраст дочери **В зависимости от него будут отличаться заданные вопросы и голос**",
        ("Вика - 10 лет", "Света - 14 лет", "Ника - 18 лет"),
    )
    noise_levels = ["Низкий", "Средний", "Высокий", "Очень высокий"]
    selected_noise_range = st.select_slider(
        "**Выберете уровень шума на фоне диалога**",
        options=noise_levels,
    )

    if selected_noise_range == "Низкий":
        noise_level = 16
    elif selected_noise_range == "Средний":
        noise_level = 12
    elif selected_noise_range == "Высокий":
        noise_level = 7
    else:
        noise_level = 3

    if "messages" not in st.session_state:
        st.session_state.messages = []
    if "file_processed" not in st.session_state:
        st.session_state.file_processed = False

    uploaded_file = st.file_uploader("Загрузите файл PDF или TXT", type=["pdf", "txt"])
    for message in st.session_state.messages:
        display_message(message['phrase'], message['role'].value)
        
    if uploaded_file and not st.session_state.file_processed:
        user_query = load_and_extract_text(uploaded_file)
        # print(type(user_query))
        # print(user_query)
        roter_result = get_router_result(user_query)
        if roter_result == 0:
            st.text(PromptClass.DEFAULT_ANSWER)
        else:
            response = get_dialogue(user_query, n=int(option_age.split()[-2])).to_json()
            # print(response)
            st.session_state.messages = response
            tts = Text2Speech(response, age=int(option_age.split()[-2]))
            tts.get_part()
            tts.combine()
            tts.post_processing(
                'audios/dialogue.wav',
                noise_level
            )
            st.rerun()

    if user_query := st.chat_input(placeholder="Answer your question"):
        roter_result = get_router_result(user_query)
        print(user_query)
        if roter_result == 0:
            st.text(PromptClass.DEFAULT_ANSWER)
        else:
            response = get_dialogue(user_query, n=int(option_age.split()[-2])).to_json()
            print(response)
            st.session_state.messages = response
            tts = Text2Speech(response, age=int(option_age.split()[-2]))
            tts.get_part()
            tts.combine()
            tts.post_processing(
                'audios/dialogue.wav',
                noise_level
            )
            st.rerun()

    if len(st.session_state.messages) > 0:
        audio_file = open('audios/final_output.wav', 'rb')
        audio_bytes = audio_file.read()
        st.audio(audio_bytes, format='audio/wav')

        
if __name__ == "__main__":
    main()
