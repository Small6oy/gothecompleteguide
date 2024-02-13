package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	//Declaring Maps
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.co",
	}

	fmt.Println("websites", websites)
	fmt.Println("Google", websites["Google"])

	// Insert
	websites["Linked"] = "https://linkedIn.com"

	// Delete
	delete(websites, "Google")

	// Update
	websites["Amazon Web Services"] = "https://aws.com"

	// Using Make
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["js"] = 5.7
	courseRatings.output()

	for index, value := range courseRatings {
		fmt.Println("Index", index, "Value", value)
	}

}
