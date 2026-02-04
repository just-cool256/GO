package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
	// to have different keys here in the output json than we have here in this struct
	// And this can be achieved by using struct tags - which is a feature built into Go
	// Struct tags are essentially metadata that you can add to your struct fields, and that metadata can then be picked up by any code or package that cares about it.
	// And it turns out the json package does care about such metadata. If it's defined, it will be considered.
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	// for json Content, it would be more comman to have lower case keys

	json, err := json.Marshal(todo)
	// Marshal function will convert data to JSON
	// this json package with its Marshal function will actually only extract and convert struct data that's made publicly available

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
	// WriteFile also returns error
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid input.")
	}
	return Todo{
		Text: text,
	}, nil
}
