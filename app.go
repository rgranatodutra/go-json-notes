package main

import "go-json-notes/notation"

func main() {
	nt := notation.New("Teste", "Olá mundo!")
	nt.ChangeTitle("Hello World")
	nt.Edit("This is my notation... the world is fantastic!")
	nt.Print()
}
