package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wfd/recipe"

	"golang.org/x/net/html"
)

func ParseRecipeFromURL(url string) (*recipe.Recipe, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	// parse recipe from body
	// search for application/ld+json
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	recipeJson := getRecipeJson(doc)
	recipe, err := parseRecipe([]byte(recipeJson))
	if err != nil {
		log.Fatalf("Failed to parse recipe: %v", err)
	}
	return recipe.ToRecipe(), nil
}

func getRecipeJson(doc *html.Node) string {
	// search for application/ld+json inside head
	head := findHead(doc)
	for c := range head.ChildNodes() {
		if c.Type == html.ElementNode && c.Data == "script" {
			for _, attr := range c.Attr {
				if attr.Key == "type" && attr.Val == "application/ld+json" {
					return c.FirstChild.Data
				}
			}
		}
	}

	return ""
}

// iterate through descendants to find head
func findHead(doc *html.Node) *html.Node {
	for c := range doc.Descendants() {
		if c.Type == html.ElementNode && c.Data == "head" {
			return c
		}
	}
	return nil
}

func parseRecipe(recipeJson []byte) (recipe.RecipeJson, error) {
	// try direct parsing
	// json is not an array
	var rec recipe.RecipeJson
	if err := json.Unmarshal(recipeJson, &rec); err == nil {
		return rec, nil
	}
	// json is an array
	var recipes []recipe.RecipeJson
	if err := json.Unmarshal(recipeJson, &recipes); err == nil {
		return recipes[0], nil
	}

	var rawJson map[string]any
	if err := json.Unmarshal(recipeJson, &rawJson); err != nil {
		return recipe.RecipeJson{}, fmt.Errorf("unmarshalling recipe: %w", err)
	}

	// If not, see if we have @graph. If so, iterate over @graph and check for @type
	if graph, ok := rawJson["@graph"].([]any); ok {
		for _, item := range graph {
			if v, ok := item.(map[string]any)["@type"]; ok {
				if v == "Recipe" {
					raw, err := json.Marshal(item)
					if err != nil {
						return recipe.RecipeJson{}, err
					}
					var rec recipe.RecipeJson
					if err := json.Unmarshal(raw, &rec); err != nil {
						return recipe.RecipeJson{}, err
					}
					return rec, nil
				}
			}
		}
	}
	return recipes[0], nil
}
