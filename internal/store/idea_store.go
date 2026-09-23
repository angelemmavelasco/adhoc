package store

import (
	"adhoc/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store interface {
	GetAll() ([]*models.Idea, error)
	GetByID(id int) (*models.Idea, error)
	Create(idea *models.Idea) (*models.Idea, error)
	Update(id int, idea *models.Idea) (*models.Idea, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*models.Idea, error) {

	return []*models.Idea{}, nil
}
func (s *store) GetByID(id int) (*models.Idea, error) {
	return &models.Idea{}, nil
}
func (s *store) Create(idea *models.Idea) (*models.Idea, error) {
	if idea.Title == "" || idea.Content == "" {
		return nil, fmt.Errorf("title or content is required")
	}
	idea.CreatedAt = time.Now()
	idea.UpdatedAt = time.Now()

	query := `
		INSERT INTO idea (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := s.db.ExecContext(ctx, query); err != nil {
		_ = s.db.Close()
		return nil, fmt.Errorf("Error while initializing tables: %w", err)
	}
	defer cancel()
	return &models.Idea{}, nil
}
func (s *store) Update(id int, idea *models.Idea) (*models.Idea, error) {
	return &models.Idea{}, nil
}
func (s *store) Delete(id int) error {
	return nil
}
