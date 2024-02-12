package main

import (
	"fmt"
)

func test(revenue, expenses, taxRate float64) (float64, float64, float64) {

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue, expenses, taxRate = test(revenue, expenses, taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Printf(`
	EBT: %.1f
	Profit: %.1f
	Ratio: %.1f`, earningsBeforeTax, profit, ratio)
}
