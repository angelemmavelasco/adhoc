package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("We are AdHoc\n")

	fmt.Println("What are we gonna do today?")
	color.Blue("[1] Crate new idea")
	color.Blue("[2] View repository")
	color.Blue("[3] How to use")

	var action string
	fmt.Scanln(&action)
	fmt.Printf("Your choice: %s", action)
}
