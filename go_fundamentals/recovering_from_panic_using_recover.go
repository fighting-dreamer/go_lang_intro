package main

import (
	"fmt"
)

func PanicRecover () {
	defer fmt.Println("Deferred call -- 1")
	defer func () {
		fmt.Println("deferred call -- 2")
		if e := recover(); e != nil {
			fmt.Println("Recovered With : ", e)
		}
	}()
	defer fmt.Println("deferred call -- 3") // let`s see if it be called or not ??

	panic("just panicing for sake of understanding recover")
	fmt.Println("This will never be called")
}

func main () {
	fmt.Println("Starting to panic..")
	PanicRecover()
	fmt.Println("this will now be printed !!")
}


