package main

import (
	"adhoc/internal/clients"
	"adhoc/internal/database"
	"adhoc/internal/handlers"
	"adhoc/internal/routes"
	"adhoc/internal/services"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error: no se pudo cargar el archivo .env: ", err)
	}

	//database connection
	dbConn, err := sql.Open("sqlite", "adhoc.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("error while connecting database %v", err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("error while checking connection to database %v", err)
	}

	queries := database.New(dbConn)

	//minimaxApiKey for llm model
	minimaxApiKey := os.Getenv("MINIMAX_API_KEY")
	if minimaxApiKey == "" {
		log.Fatal("ERROR: La variable MINIMAX_API_KEY está vacía o no se cargó")
	}
	aiClient := clients.NewMiniMaxClient(minimaxApiKey)

	//ideaService to receive ideas and process them by the llm
	ideaService := services.NewIdeaService(aiClient, queries)
	ideaHandler := handlers.NewIdeaHandler(ideaService)

	keywordService := services.NewKeywordService(queries)
	keywordHandler := handlers.NewKeywordHandler(keywordService)

	router := routes.SetupRoutes(ideaHandler, keywordHandler)

	log.Println("Listening on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error while listening on port 8080 %v", err)
	}
}
