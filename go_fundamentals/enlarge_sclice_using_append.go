package main

import (
	"fmt"
)

func main () {
	x := []int{} // right now its "nil"
	x = append(x, 10) // but you can do this without even initializaton of slice
	fmt.Println(x)

	x = append(x, 10,20,30,40)
	fmt.Println(x)

}
