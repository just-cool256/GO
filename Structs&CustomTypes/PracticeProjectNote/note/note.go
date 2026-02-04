package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	// to have different keys here in the output json than we have here in this struct
	// And this can be achieved by using struct tags - which is a feature built into Go
	// Struct tags are essentially metadata that you can add to your struct fields, and that metadata can then be picked up by any code or package that cares about it.
	// And it turns out the json package does care about such metadata. If it's defined, it will be considered.
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	// for json Content, it would be more comman to have lower case keys

	json, err := json.Marshal(note)
	// Marshal function will convert data to JSON
	// this json package with its Marshal function will actually only extract and convert struct data that's made publicly available

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
	// WriteFile also returns error
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
