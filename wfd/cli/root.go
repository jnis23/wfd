package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/jnis23/wfd/app"

	"github.com/spf13/cobra"
)

var url string
var output bool

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "The URL of the webpage to parse")
	rootCmd.Flags().BoolVarP(&output, "output", "o", false, "Output the recipe to a file")
}

var rootCmd = &cobra.Command{
	Use:   "recipe",
	Short: "A CLI tool to parse recipes from a webpage",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			cmd.Help()
			return
		}

		wfd, err := app.NewWfd()
		if err != nil {
			panic(err)
		}

		recipe, err := wfd.ParseRecipeFromURL(url)
		if err != nil {
			log.Fatalf("failed to parse recipe : %s", errors.Unwrap(err))
		}

		fmt.Printf("%+v", recipe)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute root command: %v", err)
	}
}
