package routes

import (
	"adhoc/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(ideaHandler *handlers.IdeaHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v-beta")
	{
		v1.POST("/ideas", ideaHandler.NewIdea)
		v1.GET("/ideas", ideaHandler.ReadIdeas)
	}
	return r
}
