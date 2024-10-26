import fitz

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
