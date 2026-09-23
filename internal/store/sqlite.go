package store

// this file handles db connection
import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDatabase() error {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	query := `
	create table if not exists idea (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    title VARCHAR(255) NOT NULL,
	    content VARCHAR(255) NOT NULL,
	    created_at DATETIME,
	    updated_at DATETIME
	)
	`
	if _, err := db.ExecContext(context.Background(), query); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
