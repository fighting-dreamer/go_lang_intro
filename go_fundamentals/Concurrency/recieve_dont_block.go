package main

import "fmt"

func main() {
	// Data is irrelevant
	done := make(chan int)

	go func() {
		fmt.Println("goroutine message")

		// Just send a signal "I'm done"
		close(done)
	}()

	fmt.Println("main function message")
	if nil == done {
		fmt.Println("it`s a nil channel")
	}
	if <-done == 0 {
		fmt.Println(<-done)
	}
}
