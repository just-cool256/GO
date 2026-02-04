package main

import "fmt"

// You can also use the type keyword to assign an alias to other built in types
type str string

// You can use your custom type to assign methods to it
// We couldn't use the built-in string type to assign methods to it directly.
func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Bharti"
	name.log()
}