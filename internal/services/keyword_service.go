package services

import (
	"adhoc/internal/models"
	"errors"
	"strconv"
	"strings"
)

type KeywordService interface {
	ReadKeywords() ([]*models.Keyword, error)
	NewKeyword(keyword models.Keyword) (*models.Keyword, error)
}

type keywordService struct{}

func NewKeywordService() KeywordService {
	return &keywordService{}
}

func (s *keywordService) ReadKeywords() ([]*models.Keyword, error) {
	iter := 10
	keywords := make([]*models.Keyword, 0, iter)
	for id := range iter {
		idString := strconv.Itoa(id)
		keyword := models.Keyword{
			ID:      id,
			Keyword: "keyword" + idString,
		}
		keywords = append(keywords, &keyword)
	}
	return keywords, nil
}

func (s *keywordService) NewKeyword(keyword models.Keyword) (*models.Keyword, error) {
	if len(strings.TrimSpace(keyword.Keyword)) <= 0 {
		return nil, errors.New("el nombre del t2ermino debe tener longitud de caracteres mayor a 0")
	}
	return &keyword, nil
}
