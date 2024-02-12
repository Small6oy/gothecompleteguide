package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)

	resultFloat := add(1.2, 2.4)
	fmt.Println(resultFloat)

	resultString := add("1", "2")
	fmt.Println(resultString)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
