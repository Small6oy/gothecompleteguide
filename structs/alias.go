package main

import "fmt"

type strAlias string

func (text strAlias) log() {
	fmt.Println(text)
}

func main() {
	var name strAlias = "Max"
	name.log()
}
