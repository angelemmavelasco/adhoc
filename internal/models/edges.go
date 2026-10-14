package models

type Edges struct {
	ID         int `json:"id"`
	SourceIdea int `json:"source_idea"`
	TargetIdea int `json:"target_idea"`
	Weight     int `json:"weight"`
}
