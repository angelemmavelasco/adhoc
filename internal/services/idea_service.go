package services

import (
	"adhoc/internal/models"
	"adhoc/internal/store"
	"strings"
	"time"
)

type IdeaService struct {
	store store.IdeaStore
}

func NewIdeaService(store store.IdeaStore) *IdeaService {
	return &IdeaService{store: store}
}

func (svc *IdeaService) IngestIdea(title, content string, rawKeywords []string) (*models.Idea, error) {
	keywords := make([]string, len(rawKeywords))
	for _, kw := range rawKeywords {
		cleaned := strings.ToLower(strings.TrimSpace(kw))
		if cleaned != "" {
			keywords = append(keywords, cleaned)
		}
	}

	now := time.Now()

	idea := &models.Idea{
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createIdea, err := svc.store.Create(idea)
	if err != nil {
		return nil, err
	}
	return createIdea, nil

}
