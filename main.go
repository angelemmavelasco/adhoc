package main

import (
	"adhoc/internal/handlers"
	"adhoc/internal/routes"
	"adhoc/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello world, %s", name)
}

//func getIdeas(c *gin.Context) {
//	ideas := []gin.H{
//		{"id": 1, "title": "my interests", "content": "i like machine learning for biology"},
//		{"id": 2, "title": "available careers", "content": "studies in biology, statistics"},
//	}
//	c.JSON(http.StatusOK, ideas)
//}
//
//func newIdea(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"method": "you're posting",
//	})
//}

//func putting(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"method": "you're putting",
//	})
//}
//
//func deleting(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"method": "you're deleting",
//	})
//}

func main() {
	fmt.Println("Hello")

	ideaService := services.NewIdeaService()
	ideaHandler := handlers.NewIdeaHandler(ideaService)
	router := routes.SetupRoutes(ideaHandler)

	//router.GET("/hello-world/:name", helloWorld)
	//router.POST("/new-idea", handlers.IdeaHandler{})
	//router.GET("/get-ideas", getIdeas)
	//router.PUT("/somePut", putting)
	//router.DELETE("someDelete", deleting)

	err := router.Run()
	if err != nil {
		return
	}
}
