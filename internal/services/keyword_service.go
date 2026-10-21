package services

import (
	"adhoc/internal/database"
	"adhoc/internal/models"
	"context"
	"errors"
	"fmt"
	"strings"
)

type KeywordService interface {
	ReadKeywords(ctx context.Context) ([]*models.Keyword, error)
	NewKeyword(ctx context.Context, keyword models.Keyword) (*models.Keyword, error)
}

type keywordService struct {
	queries *database.Queries
}

func NewKeywordService(queries *database.Queries) KeywordService {
	return &keywordService{
		queries: queries,
	}
}

func (s *keywordService) ReadKeywords(ctx context.Context) ([]*models.Keyword, error) {
	dbKeywords, err := s.queries.ReadKeywords(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing keywords %v", err)
	}
	kws := make([]*models.Keyword, 0, len(dbKeywords))
	for _, kw := range dbKeywords {
		kws = append(kws, &models.Keyword{
			ID:      int(kw.ID),
			Keyword: kw.Keyword,
		})
	}
	return kws, nil
}

func (s *keywordService) NewKeyword(ctx context.Context, keyword models.Keyword) (*models.Keyword, error) {
	if len(strings.TrimSpace(keyword.Keyword)) <= 0 {
		return nil, errors.New("the number of chars to create a keyword must be greater than zero")
	}

	dbKW, err := s.queries.NewKeyword(ctx, keyword.Keyword)
	if err != nil {
		return nil, fmt.Errorf("error creating keyword %v", err)
	}

	return &models.Keyword{
		ID:      int(dbKW.ID),
		Keyword: dbKW.Keyword,
	}, nil
}
