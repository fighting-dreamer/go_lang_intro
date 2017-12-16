package main


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printCounts(label string, counter chan int) {
	wg.Add(1)
	defer wg.Done()

	for val := range counter {
		fmt.Printf("%s : %d\n", label, val)
		if val == 10 {
			fmt.Printf("Channel is closed by %s\n", label)
			close(counter)
			return
		}
		val++
		//Send the value via counter channel
		counter <- val
	}
}

func main() {
	fmt.Println("Starting the program")
	var counter chan int
	counter = make(chan int)

	go printCounts("Go-routine-1", counter)
	go printCounts("Go-routine-2", counter)

	counter <- 1
	wg.Wait()
	fmt.Println("Terminating the program")
}