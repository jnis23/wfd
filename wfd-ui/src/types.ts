type Recipe = {
    Name: string;   
    Description: string;
    Ingredients: Ingredient[];
    Instructions: Instruction[];
}

type Ingredient = {
    name: string;
    quantity: string;
    unit: string;
}

type Instruction = {
    name: string;
    text: string;
}

export type { Recipe, Ingredient, Instruction };