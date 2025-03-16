package cli

import (
	"log"
	"wfd/app"
	"wfd/server"

	"github.com/spf13/cobra"
)

var addr string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the recipe server",
	Run: func(cmd *cobra.Command, args []string) {
		wfd, err := app.NewWfd()
		if err != nil {
			log.Fatalf("Failed to create wfd: %v", err)
		}
		server, err := server.NewWebServer(addr, wfd)
		if err != nil {
			log.Fatalf("Failed to create web server: %v", err)
		}
		if err := server.Serve(); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "The address to serve the web server on")
	rootCmd.AddCommand(serveCmd)
}
