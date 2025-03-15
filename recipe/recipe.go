package recipe

import (
	"encoding/json"
	"fmt"
)

// Intermediate struct to help with parsing
type RecipeJson struct {
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	PrepTime     string        `json:"prepTime,omitempty"`
	CookTime     string        `json:"cookTime,omitempty"`
	RecipeYield  Yield         `json:"recipeYield,omitempty"`
	Ingredients  []string      `json:"recipeIngredient"`
	Instructions []Instruction `json:"recipeInstructions"`
}

func (rj *RecipeJson) ToRecipe() *Recipe {
	ingredients := make([]Ingredient, len(rj.Ingredients))
	for i, ingredient := range rj.Ingredients {
		// TODO: Parse the quantity from the ingredient.
		ingredients[i] = Ingredient{
			Name: ingredient,
		}
	}
	return &Recipe{
		Name:         rj.Name,
		Description:  rj.Description,
		Ingredients:  ingredients,
		Instructions: rj.Instructions,
	}
}

type Recipe struct {
	Name         string
	Description  string
	Ingredients  []Ingredient
	Instructions []Instruction
}

func (r *Recipe) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Instruction struct {
	Name string `json:"name"`
	Step string `json:"text"`
}

type Yield struct {
	Vals []string
}

func (yield *Yield) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		yield.Vals = []string{s}
		return nil
	}

	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		yield.Vals = arr
		return nil
	}

	return fmt.Errorf("recipeYield must be either string or string array")
}
