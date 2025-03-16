package server

import (
	"log/slog"
	"net/http"
	"os"
	"wfd/app"
)

type WfdWebServer struct {
	wfd    *app.Wfd
	addr   string
	logger *slog.Logger
}

func NewWebServer(addr string, wfd *app.Wfd) (*WfdWebServer, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &WfdWebServer{
		wfd:    wfd,
		addr:   addr,
		logger: logger,
	}, nil
}

func (wws *WfdWebServer) Serve() error {
	mux := http.NewServeMux()
	wws.RegisterHandlers(mux)

	return http.ListenAndServe(wws.addr, mux)
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func (wws *WfdWebServer) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /parse", cors(wws.ParseRecipe))
}

// Parses url from query param
func (wws *WfdWebServer) ParseRecipe(w http.ResponseWriter, r *http.Request) {
	qparams := r.URL.Query()
	url := qparams.Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must pass a url query parameter."))
		wws.logger.Warn("Missing url parameters.")
		return
	}

	recipe, err := wws.wfd.ParseRecipeFromURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed parsing recipe from url."))
		wws.logger.Error("failed parsing recipe", slog.Any("error", err))
		return
	}

	recipeJson, err := recipe.ToJson()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error marshalling recipe JSON."))
		return
	}

	// return the recipe as JSON
	w.Write(recipeJson)
}
