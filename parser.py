# This model is responsible for parsing the recipes out of the cleaned HTML
from langchain_ollama import ChatOllama
from langchain.prompts import PromptTemplate
from langchain.output_parsers import PydanticOutputParser
from recipe import Recipe

def parse_recipes(cleaned_html) -> Recipe:
    model = ChatOllama(model="llama3.3")
    prompt = PromptTemplate.from_template("""
        Parse the recipe out of this HTML file. The output should be a JSON object that includes the recipe with the following fields:
        - name: The name of the recipe
        - ingredients: A list of ingredients with the following fields:
            - name: The name of the ingredient
            - quantity: The quantity of the ingredient
            - unit: The unit of the ingredient
        - instructions: A list of instructions with the following fields:
            - step: The step of the instruction
            - duration: The duration of the instruction

        Here is the HTML file:
        {cleaned_html}
    """)
    parser = PydanticOutputParser(pydantic_object=Recipe)
    chain = prompt | model | parser
    return chain.invoke({"input": cleaned_html})
