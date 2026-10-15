package handlers

import (
	"adhoc/internal/models"
	"adhoc/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IdeaHandler struct {
	service services.IdeaService
}

func NewIdeaHandler(service services.IdeaService) *IdeaHandler {
	return &IdeaHandler{service: service}
}

func (h *IdeaHandler) NewIdea(c *gin.Context) {
	var input models.Idea
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}

	createdIdea, err := h.service.NewIdea(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Idea created successfuly",
		"data":    createdIdea,
	})
}

func (h *IdeaHandler) ReadIdeas(c *gin.Context) {
	ideas, err := h.service.ReadIdeas()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Information retrieved successfuly",
		"data":    ideas,
	})
}
