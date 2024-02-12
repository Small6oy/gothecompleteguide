package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalance = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	balance, err := fileops.GetFloatFromFile(accountBalance)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("_________")
		// panic("Can't continue, sorry.")
	}

	for {
		option := menu()

		switch option {
		case 1:
			checkBalance(balance)
		case 2:
			balance = depositMoney(balance)
		case 3:
			balance = withdrawMoney(balance)
		case 4:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Please select an option between 1 - 4")
		}
	}
}

func checkBalance(balance float64) {
	fmt.Printf("Your Current Balance is %.2f\n", balance)
}

func depositMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Amount : ")
	fmt.Scan(&amount)
	fmt.Println("")

	balance += amount
	fileops.StoreFloatInFile(accountBalance, balance)
	checkBalance(balance)
	return balance
}

func withdrawMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Amount : ")
	fmt.Scan(&amount)
	fmt.Println("")

	if amount > balance {
		fmt.Printf("The amount your requested to withdraw exceeds your current funds %.2f\n", balance)
		fmt.Printf("Current balance : %.2f\nWithdraw amount: %.2f\n", balance, amount)
		fmt.Printf("Your request has been refused")
		return balance
	} else {
		balance -= amount
		fileops.StoreFloatInFile(accountBalance, balance)
		checkBalance(balance)
		return balance
	}
}
