package main

import (
	"adhoc/internal/services"
	"adhoc/internal/store"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("We are AdHoc")
	fmt.Println("We are setting all for you...")

	color.Green("Initializing ideas")
	db, err := store.InitDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	ideaStore := store.New(db)
	color.Green("Ready!\n")

	color.Green("Initializing functions...")
	reader := bufio.NewReader(os.Stdin)
	ideaService := services.NewIdeaService(ideaStore)

	fmt.Println("What are we gonna do today?")

	exit := false
	for exit != true {
		color.Blue("[1] Crate new idea")
		color.Blue("[2] View repository")  // check current notes. 10 by 10
		color.Blue("[3] Set your desk up") // used to define the api key if LLM model required
		color.Blue("[4] How to use")       // general support and user docs
		color.Red("[5] Exit")

		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		fmt.Println("\nYour choice: " + action)

		switch action {
		case "1":
			color.Blue("New ideas connect us; start creating a new one.\n")

			color.Blue("Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			color.Blue("Content: ")
			content, _ := reader.ReadString('\n')
			content = strings.TrimSpace(content)

			color.Blue("Keywords (separated by colons, e.g. internet, phone, web): ")
			rawKeywordsInput, _ := reader.ReadString('\n')
			rawKeywordsInput = strings.TrimSpace(rawKeywordsInput)
			var ideaKeyWords []string
			if rawKeywordsInput != "" {
				ideaKeyWords = strings.Split(rawKeywordsInput, ",")
			}

			newIdea, err := ideaService.IngestIdea(title, content, ideaKeyWords)
			if err != nil {
				log.Fatal(err.Error())
			}
			_ = newIdea
			fmt.Println("\nYour idea: " + newIdea.Title + " was saved successfully\n\n")
		case "2":
			color.Blue("This is how your mind is made...")
		case "3":
			color.Blue("Specify if you have LLM API key")
		case "4":
			color.Blue("We are here to help you.")
		case "5":
			color.Red("See you...")
			exit = true
		default:
			color.Blue("I don't know how to do that.")
		}
	}

}
