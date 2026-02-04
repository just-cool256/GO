package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// appUser = &user.User{
	// 	FirstName: firstName}
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// similar to
	// appappUser = user{
	// firstName,
	// lastName,
	// birthdate,
	// time.Now(),
	// } - if order is same as in struct definition
	// empty struct - appappUser = user{}
	// appUser = user{
	// firstName: firstName,
	// lastName:  lastName,
	// createdAt: time.Now(),
	// }

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// Scanln - to make sure that hitting the enter key ends the input. So notifies Go that you're done entering a value.
	fmt.Scanln(&value)
	return value
}
