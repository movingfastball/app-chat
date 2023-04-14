package main

import (
	"log"
	"net/http"

	"app-chat/internal/handlers"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}

func main() {
	mux := routes()
	log.Println("Starting web server on port 8080")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")
	_ = http.ListenAndServe("localhost:8080", mux)
}
