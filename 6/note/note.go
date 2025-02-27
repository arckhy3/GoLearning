package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type NoteData struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"create_at"`
}

func (note NoteData) Display() {
	fmt.Println("")
	fmt.Println("========================")
	fmt.Println("Your Note:")
	fmt.Println("========================")
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
	fmt.Println("CreateAt: ", note.CreatedAt)
	fmt.Println("========================")
}

func (note NoteData) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*NoteData, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid title and content")
	}
	return &NoteData{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
