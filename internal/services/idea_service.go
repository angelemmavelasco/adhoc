package services

import (
	"adhoc/internal/models"
	"errors"
	"strings"
)

type IdeaService interface {
	NewIdea(idea models.Idea) (*models.Idea, error)
}

//private struct to connect with interface
type ideaService struct {}

//constructor
func NewIdeaService() IdeaService {
	return &ideaService{}
}

func (s *ideaService) NewIdea(idea models.Idea) (*models.Idea, error) {
	if len(strings.TrimSpace(idea.Title)) < 3 {
		return nil, errors.New("el título debe tener al menos 3 caracteres")
	}

	return &idea, nil
}


