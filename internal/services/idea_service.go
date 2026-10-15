package services

import (
	"adhoc/internal/clients"
	"adhoc/internal/models"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const systemPrompt = `Eres un experto analizanod ideas de proyectos. 
Tu tarea es analizar todas las notas, ideas, listas de tareas y entradas que te de
el usuario y reconocer cuales son los términos clave.`

type IdeaService interface {
	NewIdea(idea models.Idea) (*models.Idea, error)
	ReadIdeas() ([]*models.Idea, error)
}

// ideaService private struct to connect with interface
type ideaService struct {
	aiClient clients.AIClient
}

// NewIdeaService constructor
func NewIdeaService(aiClient clients.AIClient) IdeaService {
	return &ideaService{
		aiClient: aiClient,
	}
}

func (s *ideaService) NewIdea(idea models.Idea) (*models.Idea, error) {
	if len(strings.TrimSpace(idea.Title)) < 3 {
		return nil, errors.New("el título debe tener al menos 3 caracteres")
	}
	aiResp, err := s.aiClient.ProcessInput(context.Background(), systemPrompt, idea.Content)
	if err != nil {
		return nil, err
	}
	fmt.Printf("AI response: %s", aiResp)
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
