package store

import (
	"adhoc/internal/models"
	"database/sql"
)

type IdeaToKeywordStore interface {
	Create(keyword *models.IdeaToKeyword) (*models.IdeaToKeyword, error)
}

type ideaToKwStore struct {
	db *sql.DB
}

func NewITKConn(db *sql.DB) IdeaToKeywordStore { return &ideaToKwStore{db: db} }

func (itk *ideaToKwStore) Create(ideaToKw *models.IdeaToKeyword) (*models.IdeaToKeyword, error) {
	return &models.IdeaToKeyword{}, nil
}
