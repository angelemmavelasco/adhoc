package models

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// An Idea struct works as a node in the graph database.
// The structure is composed of:
//   - Title: a representative short description.
//   - OriginalContent: the raw content entered by the user.
//   - KeyWords: a little refactor made by a llm, this keeps more clean and accessible idea.
type Idea struct {
	Title           string
	OriginalContent string
	KeyWords        []string
}

func ideaHash(i Idea) string {
	return i.Title
}

func CrateIdea() {
	var IdeaTitle string
	color.Blue("Title: ")
	fmt.Scanln(&IdeaTitle)

	var IdeaOriginalContent string
	color.Blue("Content: ")
	fmt.Scanln(&IdeaOriginalContent)

	var IdeaKeyWordsList string
	color.Blue("Keywords (separated by colons, e.g. phone, internet, web: ")
	fmt.Scanln(&IdeaKeyWordsList)
	IdeaKeyWords := strings.Split(IdeaKeyWordsList, ";")

	idea := Idea{IdeaTitle, IdeaOriginalContent, IdeaKeyWords}

	fmt.Println(idea)
	return
}
