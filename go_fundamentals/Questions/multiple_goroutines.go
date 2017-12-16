package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)
//Wait Group
var wg sync.WaitGroup
var cnt int = 0

var cntCh = make(chan int, 10)

func cntProvider(){
	for i:=0; i < 10; i++ {
		cntCh <- i
	}
}

func addOne() {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	//cnt += 1
	c := <- cntCh
	fmt.Printf("cnt = %d\n", c)
}

func main(){
	go cntProvider()
	for i := 0; i < 10; i++ {
		go addOne()
	}
	//Waiting for go routines to finish
	wg.Wait()
	fmt.Println("All goroutines are finished")
}
