package main

import (
	"fmt"
)
// please look into the "solving_deadlock_from_unbuffered_channels" program to know how to get around this
func main () {
	// Declare an un-buffered Channel
	counter := make(chan int)
	// Thins will create a Deadlock
	counter <- 10  // send operation to a channel from Main go-routine
	fmt.Println(<- counter) // recieve operation from a channel
	// This is c`oz, the send operation blocks the "main" while it search for other go-routines to recieve the counter`s value,
	// c`oz "none" other than "main" exist, thus It becomes a "deadlock"
}
