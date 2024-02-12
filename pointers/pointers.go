package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("AgePointer: ", agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
