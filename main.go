package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("We are AdHoc\n")

	fmt.Println("What are we gonna do today?")
	color.Blue("[1] Crate new idea")
	color.Blue("[2] View repository")  // check current notes. 10 by 10
	color.Blue("[3] Set your desk up") // used to define the api key if LLM model required
	color.Blue("[4] How to use")       // general support and user docs

	var action string
	fmt.Scanln(&action)
	fmt.Println("\nYour choice: " + action)
	switch action {
	case "1":
		color.Blue("New ideas connect us; start creating a new one.")
	case "2":
		color.Blue("This is how your mind is made...")
	case "3":
		color.Blue("Specify if you have LLM API key")
	case "4":
		color.Blue("We are here to help you.")
	default:
		color.Blue("I don't know how to do that.")
	}


}
