from typing import Optional
from pydantic import BaseModel


class Ingredient(BaseModel):
    name: str
    quantity: str
    unit: str

class Instruction(BaseModel):
    step: str
    duration: Optional[str]

class Recipe(BaseModel):
    name: str
    ingredients: list[Ingredient]
    instructions: list[Instruction]

