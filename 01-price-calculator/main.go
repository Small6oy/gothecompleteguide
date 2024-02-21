package main

import (
	"fmt"

	"example.com/price-calculator/tax"
)

func main() {

	menu()

	amount := getUserInput("amount: ")
	fmt.Println("amount", amount)

	taxAmount, _ := tax.GetTaxAmount(amount)
	fmt.Println("taxAmount", taxAmount)
}

func menu() {
	fmt.Println("*****************************")
	fmt.Println("WELCOME TO THE TAC CALCULATOR")
	fmt.Println("PLEASE INPUT AN AMOUNT")
	fmt.Println("*****************************")
}

func getUserInput(text string) float64 {
	fmt.Print(text)

	var input float64
	fmt.Scanln(&input)
	return input
}
