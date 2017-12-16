package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main () {
	// Add count of 2, one for each go-routine
	wg.Add(2)
	fmt.Println("Start Go-routines")
	//Launch finc as go-routines
	go MultTable()
	go AddTable()


	// Waiting for go-routines to finish
	wg.Wait()
	fmt.Println("Go routines are finished!!\nTerminating Programm!!\n")
}


func MultTable() {
	// schedule a call to waitGroup`s Done to tell the go-routine is complete
	defer wg.Done()
	for i := 1; i <= 10; i++{
		sleep := rand.Int63n(10000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Multiplication table fo i : ", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i * j)
		}
		fmt.Println()
	}
}

func AddTable() {
	// schedule a call to waitGroup`s Done to tell the go-routine is complete
	defer wg.Done()
	for i := 1; i <= 10; i++{
		sleep := rand.Int63n(10000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Addition table fo i : ", i)
		for j := 0; j <= 10; j++ {
			fmt.Printf("%d + %d = %d\n", i, j, i + j)
		}
		fmt.Println()
	}
}