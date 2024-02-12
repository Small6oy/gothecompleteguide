package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/structsPractise/note"
	"example.com/structsPractise/todo"
)

type outputtable interface {
	Display()
	Save() error
}

func main() {
	title, content := getNoteData()
	note, _ := note.New(title, content)

	text, _ := getUserInput("To Do: ")
	todo, _ := todo.New(text)

	_ = saveData(todo)
	err := saveData(note)
	if err != nil {
		return
	}
}

func saveData(data outputtable) error {

	data.Display()
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title, _ := getUserInput("Note Title: ")
	content, _ := getUserInput("Content: ")

	return title, content
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("Error Reading line")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}
