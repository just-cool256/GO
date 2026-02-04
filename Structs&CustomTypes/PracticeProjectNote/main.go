package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

// will work with single note
func main(){
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content :=getUserInput("Note content:")
	
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	
	reader := bufio.NewReader(os.Stdin) //argument describes to NewReader from which source we wanna read - command line input in this case
	// NewReader is a contructor function that will give us a value that will allow us to read user input.
	text, err := reader.ReadString('\n') // '\n' - Rune(single character)
	// text will now still contain that line break

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for Windows compatibility

	return text
}