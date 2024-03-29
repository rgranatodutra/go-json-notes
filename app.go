package main

import (
	"fmt"
	"go-json-notes/utils"
)

const notesPath string = "C:\\Users\\rgran\\OneDrive\\Área de Trabalho\\notes"
const notesExt string = ".jn"

func startApp() {
	for {
		fmt.Println("Main menu:")
		fmt.Println("1 - Create a new note")
		fmt.Println("2 - List notes")
		fmt.Println("0 - Exit program")

		chosenOption := utils.GetUserInput("Enter your choice: ")

		switch chosenOption {
		case "1":
			//createNewNote()
		case "2":
			//listSavedNotes()
		case "0":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
