package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//Package level variables, which are pointer s to log.Logger
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// initLog initializes the log.Logger Objects
func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	// Flag for defining the logging properties, to log.New
	flag := log.Ldate | log.Ltime | log.Lshortfile

	//create log.Logger Objects
	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)

}

func main() {
	initLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Trace.Println("Main Started")
	loop()
	err := errors.New("Sample Error")
	Error.Println(err.Error())
	Trace.Println("Main Completed")
}

func loop() {
	Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		Info.Println("Counter value is:", i)
	}
	Warning.Println("The counter variable is not being used")
	Trace.Println("Loop completed")
}
