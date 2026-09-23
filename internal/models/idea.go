package models

import (
	"fmt"
)

type Idea struct {
	Title               string
	OriginalContent     string
	RefactorizedContent string
}

func ideaHash(i Idea) string {
	return i.Title
}

func CrateIdea(idea Idea) {
	fmt.Println(idea)
	return
}
