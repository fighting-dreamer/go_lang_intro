package main

import (
	"fmt"
	"sync"
	//"time"
)

var wg sync.WaitGroup

func printCounts(label string, counter chan int) {
	wg.Add(1)
	defer wg.Done()
	for {
		fmt.Printf("inside %s\n", label)
		val, ok := <-counter
		if !ok {
			fmt.Printf("%s : Channel was closed!\n", label)
		}
		fmt.Printf("%s : %d\n", label, val)
		if (val == 15) {
			fmt.Printf("Channel closed from %s\n", label)
			close(counter)
			return
		}
		val++
		counter <- val
	}
}

func main() {
	var counter chan int
	counter = make(chan int)
	fmt.Println("inside main")

	go printCounts("Go-Routine-1", counter)
	go printCounts("Go-Routine-2", counter)
	//go printCounts("Go-Routine-3", counter)
	fmt.Println("sending initial value in counter")
	counter <- 1
	fmt.Println("have sent initial value")
	wg.Wait()
	fmt.Println("All Executions completed")
}
