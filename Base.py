from pydantic import BaseModel, Field
from typing import List
from enum import Enum
import json

class SpeakerRole(str, Enum):
    father = "father"
    daughter = "daughter"


class DialogEntry(BaseModel):
    role: SpeakerRole = Field(..., description="Кто говорит фразу: отец или дочь")
    phrase: str = Field(..., description="Содержание фразы")


class Dialog(BaseModel):
    phrases: List[DialogEntry] = Field(..., description="Диалог между отцом и дочерью, включающий роли и фразы")

    def to_json(self):
        dialog = [{"role": entry.role, "phrase": entry.phrase} for entry in self.phrases]
        # dialog_json = json.dumps(dialog, ensure_ascii=False, indent=2)
        return dialog
    

class RouteAnswer(BaseModel):
    result: int = Field(..., description="0 или 1, относится запрос к тематике или нет 1- относится")