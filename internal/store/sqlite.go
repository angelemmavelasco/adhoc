package store

// this file handles db connection
import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		return nil, fmt.Errorf("Error while opening database: %w", err)
	}
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Error while activating foreign keys: %w", err)
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS idea (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS keyword (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		keyword TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS idea_keyword (
		idea_id INTEGER NOT NULL,
		keyword_id INTEGER NOT NULL,
		PRIMARY KEY (idea_id, keyword_id),
		FOREIGN KEY (idea_id) REFERENCES idea(id) ON DELETE CASCADE,
		FOREIGN KEY (keyword_id) REFERENCES keyword(id) ON DELETE CASCADE
	);
	CREATE VIEW IF NOT EXISTS edges AS
	SELECT 
		ROW_NUMBER() OVER () AS id,
		nt1.idea_id AS origin_idea_id,
		nt2.idea_id AS destine_idea_id,
		COUNT(*) AS weight
	FROM idea_keyword nt1
	JOIN idea_keyword nt2 
		ON nt1.keyword_id = nt2.keyword_id 
	   AND nt1.idea_id != nt2.idea_id
	GROUP BY nt1.idea_id, nt2.idea_id;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.ExecContext(ctx, query); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Error while initializing tables: %w", err)
	}
	return db, nil
}
