package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("")
	fmt.Println("========================")
	fmt.Println("Your To Do List:")
	fmt.Println("========================")
	fmt.Println("Text: ", todo.Text)
	fmt.Println("========================")
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid todo text")
	}
	return &Todo{
		Text: text,
	}, nil
}
