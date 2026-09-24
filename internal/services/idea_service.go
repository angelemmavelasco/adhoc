package services

import (
	"adhoc/internal/models"
	"adhoc/internal/store"
	"strings"
	"time"

	"github.com/fatih/color"
)

type IdeaService struct {
	ideaStore    store.IdeaStore
	keywordStore store.KeywordStore
}

func NewIdeaService(ideaStore store.IdeaStore, keywordStore store.KeywordStore) *IdeaService {
	return &IdeaService{ideaStore: ideaStore, keywordStore: keywordStore}
}

// IngestIdea creates and persists a new idea, normalizing raw keywords and assigning timestamps.
func (svc *IdeaService) IngestIdea(title, content string, rawKeywords []string) (*models.Idea, error) {
	//set timestamp
	now := time.Now()
	//initialize an empty slice from 0 to max len raw kw capacity
	keywords := make([]string, 0, len(rawKeywords))

	//iterates over the raw kw in order to append them
	for _, kw := range rawKeywords {
		//basic cleaning by converting them to lower and removing extra spaces
		cleaned := strings.ToLower(strings.TrimSpace(kw))
		//if not exists, nothing to append
		if cleaned != "" {
			keywords = append(keywords, cleaned)
		}

		cKeyword := &models.Keyword{
			Keyword:   cleaned,
			CreatedAt: now,
			UpdatedAt: now,
		}
		_, err := svc.keywordStore.Create(cKeyword)
		if err != nil {
			return nil, err
		}
		color.Yellow("	Attached keyword %v", cleaned)
	}

	idea := &models.Idea{
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createIdea, err := svc.ideaStore.Create(idea)
	if err != nil {
		return nil, err
	}
	return createIdea, nil

}
