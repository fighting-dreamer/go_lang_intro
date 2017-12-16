package main

import (
	"fmt"
)
var s int = 0
func main() {
	fmt.Printf("Go routine : %d\n", s)
	go main()
	s += 1
	if (s == 10) {
		return
	}
	counter := make(chan int)
	counter <- 1
	fmt.Println(<-counter)
}
