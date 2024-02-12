package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are requireed")
	}

	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Admin",
			lastName:  "Admin",
			birthDate: "______",
			createdAt: time.Now(),
		},
	}
}
