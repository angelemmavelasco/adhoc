package models

type IdeaKeyword struct {
	ID        int `json:"id"`
	IdeaID    int `json:"idea_id"`
	KeywordID int `json:"keyword_id"`
}
