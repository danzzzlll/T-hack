import streamlit as st
from openai import OpenAI
from utils import load_and_extract_text
from llm import get_dialogue

st.markdown(
    """
    <style>
    .message-container {
        display: flex;
        align-items: flex-start;
        margin: 10px 0;
    }
    .avatar {
        width: 40px; /* Adjust avatar size */
        height: 40px;
        border-radius: 50%;
        margin-right: 10px;
    }
    .message {
        max-width: 80%;
        padding: 10px 15px;
        border-radius: 10px; /* Rounded corners */
        border: 1px solid #ccc;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    }
    .message.user {
        background-color: #e1f5fe;
        border-top-left-radius: 0; /* Different border-radius for user */
    }
    .message.assistant {
        background-color: #e8f5e9;
        border-top-right-radius: 0; /* Different border-radius for assistant */
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
    if role != "Дочь":
        avatar_url = "https://png.pngtree.com/png-clipart/20230824/original/pngtree-smile-father-face-expression-picture-image_8446667.png"
    else:
        avatar_url = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5I45WuCMSz3Q2q0bHR28Fx5rSmMQdnCRcOw&s"
    st.markdown(
        f"""
        <div class="message-container">
            <img src="{avatar_url}" class="avatar" alt="{role} avatar"/>
            <div class="message {role}">{message}</div>
        </div>
        """,
        unsafe_allow_html=True
    )
    # st.markdown('<div class="separator"></div>', unsafe_allow_html=True)


def main():
    st.title("Диалог по статье Отца и Дочери")

    if "messages" not in st.session_state:
        st.session_state.messages = []

    uploaded_file = st.file_uploader("Загрузите файл PDF или TXT", type=["pdf", "txt"])
    for message in st.session_state.messages:
        display_message(message['phrase'], message['role'].value)
        
    if uploaded_file:
        file_text = load_and_extract_text(uploaded_file)
        response = get_dialogue(file_text).to_json()
        st.session_state.messages = response
        del uploaded_file
        # st.rerun()

    if user_query := st.chat_input(placeholder="Answer your question"):
        response = get_dialogue(user_query)
        st.session_state.messages = response.to_json()
        st.rerun()

        
if __name__ == "__main__":
    main()