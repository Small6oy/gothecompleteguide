package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name:")
	lastName := getUserData("Please enter your last name:")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY):")

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("An Error Occurred")
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser := user.NewAdmin("gggndlovu@gmail.com", "test")
	fmt.Println(adminUser.FirstName)
	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Printf("%v ", promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
