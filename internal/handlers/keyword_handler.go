package handlers

import (
	"adhoc/internal/models"
	"adhoc/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeywordHandler struct {
	service services.KeywordService
}

func NewKeywordHandler(service services.KeywordService) *KeywordHandler {
	return &KeywordHandler{service: service}
}

func (h *KeywordHandler) NewKeyword(c *gin.Context) {
	var input models.Keyword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}

	createdKeyword, err := h.service.NewKeyword(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Keyword was created successfuly",
		"data":    createdKeyword,
	})
}
