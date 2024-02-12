package main

import (
	"fmt"
)

func menu() int {
	var option int

	fmt.Println("")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Println("")
	fmt.Print("Option : ")
	fmt.Scan(&option)
	fmt.Println("")

	return option
}
