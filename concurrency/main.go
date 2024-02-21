package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool, errChan chan string) {
	errChan <- "This is an error"

	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func main() {

	done := make(chan bool)
	err := make(chan string)

	go greet("I hope I'm liking this", done, err)
	<-done
	go slowGreet("Slowly Slowly", done, err)
	go greet("Say Whaat", done, err)
	go greet("Other Things", done, err)

}
