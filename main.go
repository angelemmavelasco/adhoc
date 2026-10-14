package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello world, %s", name)
}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "you're getting",
	})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "you're posting",
	})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "you're putting",
	})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "you're deleting",
	})
}

func main() {
	fmt.Println("Hello")

	router := gin.Default()

	router.GET("/someGet", getting)
	router.GET("/helloWorld/:name", helloWorld)

	router.POST("/somePost", posting)

	router.PUT("/somePut", putting)

	router.DELETE("someDelete", deleting)

	err := router.Run()
	if err != nil {
		return
	}
}
