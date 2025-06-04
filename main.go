package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rodabaugh/pico-link/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	platform  string
	db        *database.Queries
	siteTitle string
	subTitle  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Using enviroment variables.")
	} else {
		fmt.Println("Loaded .env file.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform != "dev" && platform != "prod" {
		log.Fatal("PLATFORM must be set to either dev or prod")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		platform: platform,
		db:       dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		MainPage(&apiCfg).Render(r.Context(), w)
	})
	mux.HandleFunc("GET /api/links", apiCfg.handlerGetAllLinks)
	mux.HandleFunc("POST /api/links", apiCfg.handlerCreateLink)

	mux.HandleFunc("GET /google", handlerRedirect)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting pico-link on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
