package main

import (
	"fmt"
	"flag"
)

func main () {
	//-h is used to get eh info of the comman line variables
	// trying to assign diff option`s values to same variable
	val1 := flag.Int("val1", 0, "just a random value")
	h := flag.String("hello", "test1", "just a string")
	h =  flag.String("hello2", "test2", "just another string")
	hh := flag.String("h", "test3", "just random string")
	flag.Parse() // to parse the values
	fmt.Println("val1:", *val1)
	fmt.Println("h:", *h) // It use the 2nd value, cool
	fmt.Println("hh:", *hh)

	//go run ./go_fundamentals/StdLibrary/experimenting_parsing_command_line.go -h= -h=dnfkjnkfnkn
	//Output :
	//val1: 0
	//h: test2
	//hh: dnfkjnkfnkn
	// it basically take the last value from the command line for a given variable
}
