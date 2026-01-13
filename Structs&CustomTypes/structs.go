package main

import (
	"fmt"
	"time"
	"errors"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Creation/Constructor function 
func newUser(firstName, lastName, birthdate string) (*user, error){
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name, and birthdate cannot be empty")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
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

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func (u user) outputUserDetails() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Account Created At:", u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// Scanln - to make sure that hitting the enter key ends the input. So notifies Go that you're done entering a value.
	fmt.Scanln(&value)
	return value
}
