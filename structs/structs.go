package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthDate string) (*user, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are requireed")
	}

	return &user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name:")
	lastName := getUserData("Please enter your last name:")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY):")

	appUser, err := newUser(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("An Error Occurred")
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Printf("%v ", promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
