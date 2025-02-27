package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputable interface {
	Saver
	Displayer
}

func main() {
	printAnything(1)
	printAnything(1.5235234)
	printAnything("etetfajfafjb qrqwrksa rqr qwr qwr ")
	title, content := getNoteData()

	notes, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(getInputData("To do text: "))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(notes)

	if err != nil {
		return
	}

	outputData(todo)
}

func printAnything(data any) {
	switch data.(type) {
	case int:
		fmt.Println("integer: ", data)
	case float64:
		fmt.Println("float64: ", data)
	case string:
		fmt.Println(data)
	}
}

func outputData(data Outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Save success")
	return err
}

func getNoteData() (string, string) {
	return getInputData("Note title: "), getInputData("Note content: ")
}

func getInputData(promt string) string {
	fmt.Print(promt)
	// fmt.Scanln(&data)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	if err != nil {
		return ""
	}
	return text
}
