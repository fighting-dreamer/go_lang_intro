package main

import "fmt"

func isPossible(n int) string {
	if (n > 2 && n%2 == 0) {
		return "YES"
	}
	return "NO"
}

func main () {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(isPossible(n))
	fmt.Println(Name)
}
