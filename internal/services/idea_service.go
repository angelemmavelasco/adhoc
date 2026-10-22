package services

import (
	"adhoc/internal/clients"
	"adhoc/internal/database"
	"adhoc/internal/models"
	"context"
	"errors"
	"fmt"
	"strings"
)

type IdeaService interface {
	NewIdea(ctx context.Context, idea models.Idea) (*models.Idea, error)
	ReadIdeas(ctx context.Context) ([]*models.Idea, error)
}

// ideaService private struct to connect with interface
type ideaService struct {
	aiClient clients.AIClient
	queries  *database.Queries
}

// NewIdeaService constructor
func NewIdeaService(aiClient clients.AIClient, queries *database.Queries) IdeaService {
	return &ideaService{
		aiClient: aiClient,
		queries:  queries,
	}
}

func (s *ideaService) NewIdea(ctx context.Context, idea models.Idea) (*models.Idea, error) {
	if len(strings.TrimSpace(idea.Title)) < 3 {
		return nil, errors.New("title must be at least 3 characters long")
	}

	dbIdea, err := s.queries.NewIdea(ctx, database.NewIdeaParams{
		Title:   idea.Title,
		Content: idea.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating idea: %v", err)
	}

	return &models.Idea{
		ID:      int(dbIdea.ID),
		Title:   dbIdea.Title,
		Content: dbIdea.Content,
	}, nil
}

func (s *ideaService) ReadIdeas(ctx context.Context) ([]*models.Idea, error) {
	dbIdeas, err := s.queries.ReadIdeas(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing ideas: %v", err)
	}
	ideas := make([]*models.Idea, 0, len(dbIdeas))
	for _, idea := range dbIdeas {
		ideas = append(ideas, &models.Idea{
			ID:      int(idea.ID),
			Title:   idea.Title,
			Content: idea.Content,
		})
	}
	return ideas, nil
}
