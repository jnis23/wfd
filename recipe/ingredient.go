// Parser for turning a string like "1 cup of flour" into an ingredient struct.
// Trying a naive regex approach for now, but could be a good candidate for an LLM.
package recipe

import (
	"regexp"
	"strings"
)

// List of common measurement units
var unitList = []string{
	"teaspoon", "teaspoons", "tbsp", "tablespoon", "tablespoons",
	"cup", "cups", "ounce", "ounces", "oz", "pound", "pounds",
	"clove", "cloves", "head", "heads", "pinch", "pinches",
	"gram", "grams", "kg", "kilogram", "kilograms", "ml", "milliliter",
	"milliliters", "liter", "liters", "quart", "quarts", "pint", "pints",
}

// Set substitute
var units = make(map[string]bool)

func init() {
	for _, unit := range unitList {
		units[unit] = true
	}
}

// Regular expression to match the ingredient format
var ingredientRegex = regexp.MustCompile(`^([\d/.\s]+)?\s*([\w-]+)?\s*(.*)$`)

// ParseIngredient extracts the quantity, unit, and ingredient from a given line
func ParseIngredient(line string) Ingredient {
	line = strings.TrimSpace(line)
	matches := ingredientRegex.FindStringSubmatch(line)

	if matches == nil {
		return Ingredient{Name: line} // If no match, return the whole line as an ingredient
	}

	quantity := strings.TrimSpace(matches[1])
	unit := strings.ToLower(strings.TrimSpace(matches[2]))
	ingredient := strings.TrimSpace(matches[3])

	// Check if the extracted unit is a known unit
	if !units[unit] {
		// If the "unit" is not recognized, it's probably part of the ingredient
		ingredient = strings.TrimSpace(unit + " " + ingredient)
		unit = ""
	}

	return Ingredient{
		Quantity: quantity,
		Unit:     unit,
		Name:     ingredient,
	}
}
