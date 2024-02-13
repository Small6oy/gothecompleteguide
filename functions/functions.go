package main

import "fmt"

type transformFn func(int) int

func main() {

	// Use functions as parameter
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// Receive functions as return
	transformFunc1 := getTransformerFunction(&numbers)
	transformFunc2 := getTransformerFunction(&moreNumbers)

	func1 := transformNumbers(&numbers, transformFunc1)
	func2 := transformNumbers(&numbers, transformFunc2)

	fmt.Println(func1)
	fmt.Println(func2)

	//Closure and Factory
	doubleTransformer := createTransformer(2)
	func3 := transformNumbers(&numbers, doubleTransformer)
	fmt.Println(func3)

	//Recursive Function
	factorialNumber := factorial(6)
	fmt.Println(factorialNumber)

	//Variadic Function
	variadicSum := sumup(1, 2, 3, 4, 5, 6)
	fmt.Println(variadicSum)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
