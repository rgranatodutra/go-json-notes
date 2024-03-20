package main

import "go-json-notes/notation"

func main() {
	nt := notation.New("Teste", "Olá mundo!")
	nt.ChangeTitle("Hello World")
	nt.ChangeText("This is my notation... the world is fantastic!")
	nt.Print()
}
