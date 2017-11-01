package main

import "fmt"

func bluesky() int {
	return 5;
}

func main () {
	 a, _ := fmt.Println("hello World!")
	fmt.Println(a) // println return length of the string/stuff it is printing
	fmt.Println("this is great!!", bluesky())
}
