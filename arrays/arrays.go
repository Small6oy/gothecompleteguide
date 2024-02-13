package main

import "fmt"

func main() {

	//Declaring Arrays
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	productNames := []string{}

	fmt.Println("price", prices[0])
	fmt.Println("prices", prices)
	fmt.Println("productNames", len(productNames), cap(productNames))

	// Get Array Item
	fmt.Println("price", prices[0])

	// Array Slice
	featuredPrices := prices[1:3]
	fmt.Println("featuredPrices", featuredPrices)
	sliceFromStartToEndElement := prices[:3]
	fmt.Println("sliceFromStartToEndElement", sliceFromStartToEndElement)
	sliceStartFromElementToEnd := prices[1:]
	fmt.Println("sliceStartFromElementToEnd", sliceStartFromElementToEnd)

	//Array Added
	productNames = append(productNames, "product1", "product2")
	fmt.Println("productNames", productNames)
	fmt.Println("productNames", len(productNames), cap(productNames))

	userNames := make([]string, 2, 5)
	userNames[0] = "Sindy"
	userNames = append(userNames, "Godwin")
	userNames = append(userNames, "Gabriel")

	fmt.Println("userNames", userNames)

	for index, value := range userNames {
		fmt.Println("Index", index, "Value", value)
	}

}
