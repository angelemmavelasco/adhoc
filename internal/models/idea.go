package models

import (
	"time"
)

// An Idea struct works as a node in the graph database.
// The structure is composed of:
//   - ID: idea id.
//   - Title: a representative short description.
//   - Content: the raw content entered by the user.
//   - KeyWords: slice which contains the keyword that belong to this idea.
type Idea struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
