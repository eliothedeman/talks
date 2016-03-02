package main

import "fmt"

func dumbFunc() {
	// 7 OMIT
	unbuffered := make(chan int)       // synchronous
	buffered := make(chan float64, 10) // asynchronous until full
	// 8 OMIT
	println(unbuffered, buffered)
}

func main() {
	// START OMIT

	// Create a channel of type string
	// 1 OMIT
	c := make(chan string)
	// 2 OMIT

	// Start a goroutine with an annonomous function that recieves from the channel
	// and prints the result
	go func() {
		// 5 OMIT
		newStr := <-c
		// 6 OMIT

		fmt.Println(newStr)
	}()

	// Send the message "hello world" on the channel
	// 3 OMIT
	c <- "hello world"
	// 4 OMIT

	// END OMIT
}
