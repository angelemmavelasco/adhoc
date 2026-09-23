package models

import "time"

type Keyword struct {
	ID        int64
	Keyword   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
