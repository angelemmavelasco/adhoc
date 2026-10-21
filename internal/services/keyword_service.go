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
	NewKeyword(keyword models.Keyword) (*models.Keyword, error)
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

func (s *keywordService) NewKeyword(keyword models.Keyword) (*models.Keyword, error) {
	if len(strings.TrimSpace(keyword.Keyword)) <= 0 {
		return nil, errors.New("el nombre del t2ermino debe tener longitud de caracteres mayor a 0")
	}
	return &keyword, nil
}
