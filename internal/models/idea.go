package models

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// An Idea struct works as a node in the graph database.
// The structure is composed of:
//   - ID: idea id.
//   - Title: a representative short description.
//   - Content: the raw content entered by the user.
//   - KeyWords: slice which contains the keyword that belong to this idea.
type Idea struct {
	ID       int64
	Title    string
	Content  string
	KeyWords []string
}

func CreateIdea() {
	var IdeaTitle string
	color.Blue("Title: ")
	fmt.Scanln(&IdeaTitle)

	var IdeaContent string
	color.Blue("Content: ")
	fmt.Scanln(&IdeaContent)

	var IdeaKeyWordsList string
	color.Blue("Keywords (separated by colons, e.g. phone, internet, web: ")
	fmt.Scanln(&IdeaKeyWordsList)
	IdeaKeyWords := strings.Split(IdeaKeyWordsList, ";")

	idea := Idea{IdeaTitle, IdeaContent, IdeaKeyWords}

	fmt.Println(idea)
	return
}
