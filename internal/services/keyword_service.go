package services

import (
	"adhoc/internal/models"
	"strconv"
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
	keyword = models.Keyword{
		ID:      1,
		Keyword: "k1",
	}
	return &keyword, nil
}
