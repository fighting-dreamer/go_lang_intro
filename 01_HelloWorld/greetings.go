package main

import "fmt"

func Greetings (prefix string, who ...string) {
	for e := range who {
		fmt.Println(prefix, " ", who[e])
	}
}
