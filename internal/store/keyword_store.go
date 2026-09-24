package store

import (
	"adhoc/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type KeywordStore interface {
	GetAll() ([]*models.Keyword, error)
	GetByID(id int) (*models.Keyword, error)
	Update(id int, keyword *models.Keyword) (*models.Keyword, error)
	Create(keyword *models.Keyword) (*models.Keyword, error)
	Delete() error
}

type keywordStore struct {
	db *sql.DB
}

func NewKeywordConn(db *sql.DB) KeywordStore { return &keywordStore{db: db} }

func (kw *keywordStore) GetAll() ([]*models.Keyword, error) {
	return []*models.Keyword{}, nil
}

func (kw *keywordStore) GetByID(id int) (*models.Keyword, error) {
	return &models.Keyword{}, nil
}

func (kw *keywordStore) Create(keyword *models.Keyword) (*models.Keyword, error) {
	if keyword.Keyword == "" {
		return nil, fmt.Errorf("keyword argument must be provided")
	}
	now := time.Now()

	query := `INSERT INTO keyword (keyword, created_at, updated_at) values (?,?,?)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := kw.db.ExecContext(ctx, query, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to insert keyword: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last inserted id: %w", err)
	}

	keyword.ID = id

	return keyword, nil
}

func (kw *keywordStore) Update(id int, keyword *models.Keyword) (*models.Keyword, error) {
	return &models.Keyword{}, nil
}

func (kw *keywordStore) Delete() error {
	return nil
}
