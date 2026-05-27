package main

import (
	"log"
	"net/http"

	"github.com/Rha02/symfonia-backend/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var PORT = ":8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	handlers.NewHandlers(handlers.NewRepository())

	router := newRouter()

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on port %s", PORT)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/stocks", handlers.Repo.GetStocks)

	return r
}
