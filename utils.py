import fitz
import streamlit as st

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


def load_and_extract_text(uploaded_file):
    file_extension = uploaded_file.name.split(".")[-1].lower()
    if file_extension == "txt":
        text = uploaded_file.read().decode("utf-8")
    elif file_extension == "pdf":
        text = ""
        with fitz.open(stream=uploaded_file.read(), filetype="pdf") as pdf_document:
            for page_num in range(pdf_document.page_count):
                page = pdf_document[page_num]
                text += page.get_text()
    return text


def get_md_style():
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