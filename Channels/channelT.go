package main

import "fmt"

func main() {

	// initializing a channel
	message := make(chan string)

	go printhello(message)

	//reading message from channel
	fmt.Println(<-message)
}

func printhello(message chan string) {

	//passing message to channel
	message <- "Hello World!! "

}
