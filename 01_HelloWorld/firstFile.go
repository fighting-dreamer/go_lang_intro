package main

import (
	"fmt"
	"github.com/fgtdreamer/go_lang_intro/02_Utilities/stringutil"
)

func bluesky() int {
	return 5
}

func main() {
	a, _ := fmt.Println("hello World!")
	fmt.Println(a) // println return length of the string/stuff it is printing
	fmt.Println("this is great!!", bluesky())
	//"trying functional paradigm in go"
	fmt.Println("this is actually working and auto imported the library!", stringutil.MyLength("this is actually working and auto imported the library!"))
	//Greetings("Good bye! ", "Nipun", "Qwerty", "Avinash")
	stringutil.Test()
	stringutil.Test2()

	stringutil.AddFavourite("it seems to recoganize")
	stringutil.PrintFavourites()
	fmt.Println(len(stringutil.GetAllFavourites()))

	stringutil.GetAllFavourites2()
}
