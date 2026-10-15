package routes

import (
	"adhoc/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(ideaHandler *handlers.IdeaHandler, keywordHandler *handlers.KeywordHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v-beta")
	{
		v1.POST("/ideas", ideaHandler.NewIdea)
		v1.GET("/ideas", ideaHandler.ReadIdeas)
		v1.POST("/keywords", keywordHandler.NewKeyword)
		v1.GET("/keywords", keywordHandler.ReadKeywords)
	}
	return r
}
