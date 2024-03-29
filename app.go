package main

import (
	"errors"
	"fmt"
	"go-json-notes/notes"
	"go-json-notes/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const notesPath string = "C:\\Users\\rgran\\OneDrive\\Área de Trabalho\\notes"
const notesExt string = ".jn"

func main() {
	startApp()
}

func startApp() {
	for {
		fmt.Println("Main menu:")
		fmt.Println("1 - Create a new note")
		fmt.Println("2 - List notes")
		fmt.Println("0 - Exit program")

		chosenOption := utils.GetUserInput("Enter your choice: ")

		switch chosenOption {
		case "1":
			createNewNote()
		case "2":
			listSavedNotes()
		case "0":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func createNewNote() {
	title := utils.GetUserInput("Enter note title: ")
	content := utils.GetUserInput("Enter note content: ")

	newNote, err := notes.New(title, content)
	utils.HandleError(err)

	err = newNote.Save(notesPath)
	utils.HandleError(err)

	fmt.Printf("Succesful created: %s\n", title)
}

func listSavedNotes() {
	var files []string

	err := filepath.Walk(notesPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), notesExt) {
			files = append(files, info.Name())
		}
		return nil
	})

	utils.HandleError(err)

	for index, file := range files {
		fmt.Printf("%d - %s\n", (index + 1), file)
	}

	selectNote(files)
}

func selectNote(files []string) {
	selectedOption := utils.GetUserInput("Type \"0\" to exit or any other number to select a note to open: ")

	if selectedOption == "0" {
		return
	}

	numberOption, err := strconv.Atoi(selectedOption)

	isValidOption := numberOption > 0 && numberOption <= len(files)

	if !isValidOption {
		utils.HandleError(errors.New("invalid option, try again"))
		return
	}

	notePath := notesPath + "/" + files[numberOption-1]
	selectedNote, err := notes.Open(notePath)
	utils.HandleError(err)

	selectedNote.Print()
}
