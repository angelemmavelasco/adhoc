package models

import "github.com/dominikbraun/graph"

type Idea struct {
	Title               string
	OriginalContent     string
	RefactorizedContent string
}

func IdeaHash(i Idea) string {
	return i.Title
}

func CrateIdea(idea Idea) (nil, err) {
	g := graph.New(IdeaHash)
	_ = g.AddVertex(london)
}
