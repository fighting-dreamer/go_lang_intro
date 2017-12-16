package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func AddTable() {
	// To call at end of the AddTable Method
	defer wg.Done()

	for i := 0; i < 10; i++ {
		//sleep := rand.Int63n(1000)
		time.Sleep(time.Millisecond )
		for j := 0; j < 1000; j++ {
			fmt.Printf("+++++++++++++")
		}
		fmt.Println()
	}

}

func MultTable() {
	// To call at end of the AddTable Method
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		//sleep := rand.Int63n(1000)
		time.Sleep(time.Millisecond)
		for j := 1; j <= 1000; j++ {
			fmt.Printf("***********")
		}
		fmt.Println()
	}

}

func main() {
	wg.Add(2)

	go AddTable()
	go MultTable()

	// Waiting for go routines to finish
	wg.Wait()
	fmt.Println("Goroutines are finished")
}
