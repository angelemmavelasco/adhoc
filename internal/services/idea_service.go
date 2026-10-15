package services

import (
	"adhoc/internal/models"
	"errors"
	"strconv"
	"strings"
)

type IdeaService interface {
	NewIdea(idea models.Idea) (*models.Idea, error)
	ReadIdeas() ([]*models.Idea, error)
}

// ideaService private struct to connect with interface
type ideaService struct{}

// NewIdeaService constructor
func NewIdeaService() IdeaService {
	return &ideaService{}
}

func (s *ideaService) NewIdea(idea models.Idea) (*models.Idea, error) {
	if len(strings.TrimSpace(idea.Title)) < 3 {
		return nil, errors.New("el título debe tener al menos 3 caracteres")
	}

	return &idea, nil
}

func (s *ideaService) ReadIdeas() ([]*models.Idea, error) {
	iter := 10
	ideas := make([]*models.Idea, 0, iter)
	for id := range iter {
		idString := strconv.Itoa(id)
		newIdea := models.Idea{
			ID:      id,
			Title:   "title" + idString,
			Content: "Generic content " + idString,
		}

		ideas = append(ideas, &newIdea)

	}
	return ideas, nil
}
