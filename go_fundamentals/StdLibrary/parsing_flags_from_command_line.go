package main

import (
	"fmt"
	"flag"
)
// In this example to log and provide file path via command line
func main() {
	//flag.Type("name", "default_value", "Description")
	fileName := flag.String("fileName", "logFile", "File Name for Log File")
	logLevel := flag.Int("loglevel", 0, "An integer for LogLevel(0-4)")
	isEnable := flag.Bool("isEnable", false, "A boolean value for enabling log options")

	var num int
	//Bind flags to variable
	flag.IntVar(&num, "num", 25, "An integer value")
	//Parse parses flag definations from arguments lists
	flag.Parse()
	//Get values from the pointers
	fmt.Println("filename :", *fileName)
	fmt.Println("logLevel :", *logLevel)
	fmt.Println("isEnable :", *isEnable)
	//Get teh value from the variable
	fmt.Println("num :", num)
	//Args return the non-flag command -line arguments
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("Non-Flag command line arguments are : ")
		for _, v := range args {
			fmt.Println(v)
		}
	}
}

