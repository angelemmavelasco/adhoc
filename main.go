package main

import (
	"adhoc/internal/clients"
	"adhoc/internal/handlers"
	"adhoc/internal/routes"
	"adhoc/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func helloWorld(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello world, %s", name)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error: no se pudo cargar el archivo .env: ", err)
	}

	//minimaxApiKey for llm model
	minimaxApiKey := os.Getenv("MINIMAX_API_KEY")
	if minimaxApiKey == "" {
		log.Fatal("ERROR: La variable MINIMAX_API_KEY está vacía o no se cargó")
	}
	aiClient := clients.NewMiniMaxClient(minimaxApiKey)

	//ideaService to receive ideas and process them by the llm
	ideaService := services.NewIdeaService(aiClient)
	ideaHandler := handlers.NewIdeaHandler(ideaService)

	keywordService := services.NewKeywordService()
	keywordHandler := handlers.NewKeywordHandler(keywordService)

	router := routes.SetupRoutes(ideaHandler, keywordHandler)

	err := router.Run()
	if err != nil {
		return
	}
}
