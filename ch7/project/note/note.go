package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var path = "ch7/project/data/"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note Titled '%v' has the following Content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_"))
	fileExt := "json"
	json_data, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(path+fileName+"."+fileExt, json_data, 0644)
}

func New(noteTitle, noteContent string) (Note, error) {
	if noteTitle == "" || noteContent == "" {
		return Note{}, errors.New("fields can not be empty")
	}
	return Note{
		Title:     noteTitle,
		Content:   noteContent,
		CreatedAt: time.Now(),
	}, nil
}
