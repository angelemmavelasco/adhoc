package services

import (
	"adhoc/internal/clients"
	"adhoc/internal/database"
	"adhoc/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

const systemPrompt = `Eres un experto analizando ideas de proyectos.
Tu tarea es extraer los conceptos o términos clave más relevantes del texto proporcionado.
Debes responder ÚNICAMENTE con un JSON válido en forma de arreglo de strings, sin texto introductorio ni explicaciones adicionales.
Ejemplo de formato:
["machine learning", "algoritmos", "optimizacion"]`

func parseKW(raw string) ([]string, error) {
	clean := strings.TrimSpace(raw)
	start := strings.Index(clean, "[")
	end := strings.LastIndex(clean, "]")
	if start != -1 && end != -1 && end > start {
		clean = clean[start : end+1]
	}
	var keywords []string
	if err := json.Unmarshal([]byte(clean), &keywords); err != nil {
		return nil, err
	}

	return keywords, nil
}

type Node struct {
	Idea     *models.Idea
	Keywords []models.Keyword
}

type NodeService interface {
	NewNode(ctx context.Context, idea models.Idea) (*Node, error)
}

type nodeService struct {
	aiClient    clients.AIClient
	ideaService IdeaService
	queries     *database.Queries
}

func NewNodeService(
	aiClient clients.AIClient,
	ideaService IdeaService,
	queries *database.Queries,
) NodeService {
	return &nodeService{
		aiClient:    aiClient,
		ideaService: ideaService,
		queries:     queries,
	}
}

func (s *nodeService) NewNode(ctx context.Context, idea models.Idea) (*Node, error) {

	//create base idea and register it
	createdIdea, err := s.ideaService.NewIdea(ctx, models.Idea{
		Title:   idea.Title,
		Content: idea.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating new idea: %w", err)
	}

	//ask llm for looking in the idea for keywords
	aiResp, err := s.aiClient.ProcessInput(ctx, systemPrompt, idea.Content)
	if err != nil {
		return nil, err
	}
	fmt.Printf("AI response: %s", aiResp)

	//parse ai resp in ["machine learning", "algoritmos", "optimizacion"] format
	parsedWords, err := parseKW(aiResp)
	if err != nil {
		fmt.Printf("Error parsing kw words: %v", err)
		return &Node{
			Idea:     createdIdea,
			Keywords: []models.Keyword{},
		}, nil
	}

	//incorporate keywords in Keyword slice
	var keywords []models.Keyword
	for _, word := range parsedWords {
		keywords = append(keywords, models.Keyword{Keyword: word})
	}

	node := &Node{
		Idea:     createdIdea,
		Keywords: keywords,
	}

	return node, nil

}
