package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main () {
	//Declaring a channel
	counter := make(chan int)

	//Launch of Go-routine with label : "Ist Go-routine"
	go PrintCounts("Ist Go-routine", counter)
	//Launch of Go-routine with label : "2nd Go-routine"
	go PrintCounts("2nd Go-routine", counter)
	fmt.Println("Communication between go-routines begins")

	counter <- 1
	// Waiting to finish!!
	fmt.Println("Waiting to finish!!")
	wg.Wait()
	fmt.Println("It completed!!")
}

func PrintCounts(label string, counter chan int) {
	defer wg.Done()
	for {
		// Recieve from a Channel
		fmt.Println("Inside ", label)
		val, ok := <- counter
		if !ok {
			fmt.Println("Channel was closed!!")
			return
		}
		fmt.Printf("label : %s , val : %d\n", label, val)
		if val == 10 {
			fmt.Println("Channel Closed from %s\n", label)
			// Close the Channel
			close(counter)

			return
		}

		val++
		// sent count back to original routine
		counter <- val
	}
}

/*
1.
Communication between go-routines begins
label : 2nd Go-routine , val : 1
Waiting to finish!!
It completed!!


2.
Communication between go-routines begins
Inside  Ist Go-routine
Inside  2nd Go-routineq
Waiting to finish!!
It completed!!
label : Ist Go-routine , val : 1
Inside  Ist Go-routine
label : 2nd Go-routine , val : 2

3.
Communication between go-routines begins
Inside  Ist Go-routine
Inside  2nd Go-routine
label : Ist Go-routine , val : 1
Inside  Ist Go-routine
label : 2nd Go-routine , val : 2
Inside  2nd Go-routine
label : Ist Go-routine , val : 3
Waiting to finish!!
Inside  Ist Go-routine
It completed!!

4.
Communication between go-routines begins
Inside  Ist Go-routine
label : Ist Go-routine , val : 1
Waiting to finish!!
It completed!!
 */