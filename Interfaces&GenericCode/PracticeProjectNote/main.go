package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// Creating interfaces
// interface is a contract - that guarantees that a certain value typically a struct, has a certain method.
// here, i wanna guarantee that whichever struct implements this interface, so whichever struct signs this contract, has a Save method.
// There is a convention in Go that if you build an interface that defines only one method, and you could have more methods, you can add as many methods signatures here as you want. But if you have an Interface that requires only one method, then your interface name is that method name plus the suffix "er". - it is a common convention in Go.
type saver interface {
	Save() error
	// Save(int, string) error
	// interfaces are not about defining the logic of a method. But instead they simply define that a certain method exists that it's there and what its name is and what its return values are. In addition, an interface can also define that a method accepts certain types of values as input parameters.
}

// An interface is a type - that means that you can use it anywhere where a type is needed.

// type displayer interface {
// Display()
// }

// type outputtable interface {
// Save() error
// Display()
// }

// type outputtable interface {
// saver
// displayer
// }

type outputtable interface {
	saver
	Display()
}

// will work with single note
func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()
	// we can pass todo and userNote as values to saveData, which wants a value of type saver bcz the above interface says that a valid value must have a Save method, which takes no arguments and returns an error.
	// err = saveData(todo)
	err = outputData(todo)

	if err != nil {
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	outputData(userNote)

	printSomething(10)
	printSomething(10.11)
	printSomething("hello")

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// it guarantee that whatever data will be, whichever concrete type of value it will be, it will be of some type that adheres to that saver interface, so that signed, that saver interface contract.
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving:", err)
		return err
	}
	fmt.Println("Saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

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

// In most Go apps, you won't have too many functions that should accept any kind of value because that can also be dangerous because you might start accepting values which you don't actually want to handle.
// any value is allowed type
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// fmt.Println("Integer:", value)
	// case float64:
	// fmt.Println("Float:", value)
	// case string:
	// fmt.Println(value)
	// }
	// switch value.(type){
	// can define now different cases for the different types you might want to handle
	// And here you therefore now don't specify values that could be stored in this variable, but instead you specify different value types, for example int
	// case int:
	// ...
	// }
	// fmt.Println(value)
	// the built-in Println function, also accepts essentially any kind of value
}

// Sometimes you wanna accept any kind of value, but still don't wanna do something for all possible kinds of values that might be arriving. So, Go gives you a special version of switch statement  - with which you can also look at the type of a value with special syntax
// switch value.(type)
