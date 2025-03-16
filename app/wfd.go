package app

import (
	"log/slog"
	"os"
	"wfd/parser"
	"wfd/recipe"
)

type Wfd struct {
	logger *slog.Logger
}

func NewWfd() (*Wfd, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Wfd{
		logger: logger,
	}, nil
}

// Extracts a recipe object from the target URL
func (wfd *Wfd) ParseRecipeFromURL(url string) (*recipe.Recipe, error) {
	return parser.ParseRecipeFromURL(url)
}
