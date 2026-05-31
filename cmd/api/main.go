package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Rha02/symfonia-backend/src/dbrepo"
	"github.com/Rha02/symfonia-backend/src/driver"
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

	dbConn := os.Getenv("DB_CONNECTION")
	if dbConn == "" {
		log.Fatal("Missing DB_CONNECTION env variable!")
	}

	db, err := driver.ConnectSQL(dbConn)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	dbRepo := dbrepo.NewPostgresRepo(db.SQL)

	handlers.NewHandlers(handlers.NewRepository(dbRepo))

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
	r.Get("/search/stocks", handlers.Repo.SearchStock)
	r.Get("/stocks/{symbol}", handlers.Repo.GetStockBySymbol)

	return r
}
