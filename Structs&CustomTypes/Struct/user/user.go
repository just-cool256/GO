package user

import (
	"errors"
	"fmt"
	"time"
)

// Uppercase starting character rule for exporting something does actually not just apply to the struct type overall, but instead also to all the fields in a struct. You can export a struct and still not export all the fields in that struct.
// If, for eg, are certain fields that should not be accessible from outside the package,
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Struct Embedding - the idea is that you build a new struct that builds up on an existing struct

type Admin struct {
	email    string
	password string
	// AUser    User
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "--/--/----",
			createdAt: time.Now(),
		},
	}
}

// Creation/Constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name, and birthdate cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Account Created At:", u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
