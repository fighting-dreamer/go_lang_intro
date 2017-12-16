package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 2, 3, 4, 5,}
	fmt.Println(x)

	langs := [4]string{0: "Go", 3: "Julia"}
	langs[1] = "Rust"
	langs[2] = "Scala"
	fmt.Println(langs)

	y := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(y)

	for _, v := range (langs) {
		fmt.Println(v)
	}

	for i := 0; i < len(langs); i++ {
		fmt.Println(langs[i])
	}
}
