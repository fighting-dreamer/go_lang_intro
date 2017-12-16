package main

import (
	"fmt"
	"sync"
	//"time"
	//"math/rand"
)
//Wait Group
var wg sync.WaitGroup
var cnt int = 0

//var cntCh = make(chan int, 10)

func cntProvider(cntCh chan int, n int){
	defer wg.Done()
	for i:=0; i < n; i++ {
		cntCh <- i
	}
}

func addOne(cntCh chan int, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		//time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		//cnt += 1
		c := <-cntCh
		fmt.Printf("cnt = %d\n", c)
	}
}

func main(){
	wg.Add(2)
	cntCh := make(chan int)
	n := 10
	go cntProvider(cntCh, n)
	go addOne(cntCh, n)

	//Waiting for go routines to finish
	wg.Wait()
	fmt.Println("All goroutines are finished")
}
