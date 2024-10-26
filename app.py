import streamlit as st
from openai import OpenAI
from utils import load_and_extract_text
from llm import get_dialogue, get_router_result
from Prompts import PromptClass
from speech import Text2Speech

st.markdown(
    """
    <style>
    .message-container {
        display: flex;
        align-items: flex-start;
        margin: 10px 0;
    }
    .avatar {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        margin: 0 10px;
    }
    .message {
        max-width: 70%;
        padding: 10px 15px;
        border-radius: 10px;
        border: 1px solid #ccc;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    }
    /* User message alignment and background color */
    .message-container.daughter {
        justify-content: flex-end; /* Align user messages to the right */
    }
    .message.daughter {
        background-color: #e1f5fe;
        border-top-left-radius: 10px;
        border-top-right-radius: 0;
    }
    /* Assistant message alignment and background color */
    .message-container.father {
        justify-content: flex-start; /* Align assistant messages to the left */
    }
    .message.father {
        background-color: #e8f5e9;
        border-top-right-radius: 10px;
        border-top-left-radius: 0;
    }
    .separator {
        height: 1px;
        background-color: #ccc;
        margin: 20px 0;
    }
    </style>
    """,
    unsafe_allow_html=True
)


def display_message(message, role):
    if role != "daughter":
        avatar_url = "https://png.pngtree.com/png-clipart/20230824/original/pngtree-smile-father-face-expression-picture-image_8446667.png"
    else:
        avatar_url = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5I45WuCMSz3Q2q0bHR28Fx5rSmMQdnCRcOw&s"
    st.markdown(
        f"""
        <div class="message-container {role}">
            <div>
                <img src="{avatar_url}" class="avatar" alt="{role} avatar"/>
            </div>
            <div class="message {role}">{message}</div>
        </div>
        """,
        unsafe_allow_html=True
    )


def main():
    st.title("Диалоги о Главном: Отец и Дочь")
    option = st.selectbox(
        "Выберете возвраст дочери? В зависимости от него будут отличаться зданные вопросы и голос",
        ("Вика - 10 лет", "Света - 14 лет", "Ника - 18 лет"),
    )

    if "messages" not in st.session_state:
        st.session_state.messages = []
    if "file_processed" not in st.session_state:
        st.session_state.file_processed = False

    uploaded_file = st.file_uploader("Загрузите файл PDF или TXT", type=["pdf", "txt"])
    for message in st.session_state.messages:
        display_message(message['phrase'], message['role'].value)
        
    if uploaded_file and not st.session_state.file_processed:
        file_text = load_and_extract_text(uploaded_file)
        roter_result = get_router_result(file_text)
        if roter_result == 0:
            st.text(PromptClass.DEFAULT_ANSWER)
        else:
            response = get_dialogue(file_text).to_json()
            st.session_state.messages = response
            st.session_state.file_processed = True
            tts = Text2Speech(response, age=int(option.split()[-2]))
            tts.get_part()
            tts.combine()
            tts.post_processing('audios/dialogue.wav')
            st.rerun()

    if user_query := st.chat_input(placeholder="Answer your question"):
        roter_result = get_router_result(user_query)
        if roter_result == 0:
            st.text(PromptClass.DEFAULT_ANSWER)
        else:
            response = get_dialogue(user_query).to_json()
            st.session_state.messages = response
            tts = Text2Speech(response, age=int(option.split()[-2]))
            tts.get_part()
            tts.combine()
            tts.post_processing('audios/dialogue.wav')
            st.rerun()

    if len(st.session_state.messages) > 0:
        audio_file = open('audios/final_output.wav', 'rb')
        audio_bytes = audio_file.read()
        st.audio(audio_bytes, format='audio/wav')

        
if __name__ == "__main__":
    main()
