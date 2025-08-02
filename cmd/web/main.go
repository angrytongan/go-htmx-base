package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	defaultPort = 8886
)

func run() error {
	app := newApplication()
	mux := chi.NewRouter()
	assetFileServer := http.FileServer(http.Dir("./assets"))

	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(5, "text/html", "text/css", "text/javascript"))
	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)
	mux.Handle("/img/*", assetFileServer)
	app.setRoutes(mux)

	server := newServer(defaultPort, mux)

	log.Printf("listening on port %d\n", defaultPort)
	err := server.ListenAndServe()

	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("run(): %v", err)
	}
}
