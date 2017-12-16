package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

const (
	//UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	TRACE                    // 1
	INFO                     // 2
	WARNING                  // 3
	ERROR                    // 4
)

// Level holds the log level.
type Level int

//Package level variables, which are pointers to log.Logger
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

//Initialize the log.Logger Objects
func initLog(traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	isFlag bool) {

	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}

	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "Error: ", flag)
}

//SetLogLevel sets the logging level
func SetLogLevel(level Level) {
	// Create an os.*File, which has implemented io.Writer interface
	f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening the log file: %s", err.Error())
	}

	//Call initLog func by specifying the logLevel
	switch level {
	case TRACE:
		initLog(f, f, f, f, true)
		return
	case INFO:
		initLog(ioutil.Discard, f, f, f, true)
		return
	case WARNING:
		initLog(ioutil.Discard, ioutil.Discard, f, f, true)
		return
	case ERROR:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, f, true)
		return
	default:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, false)
		f.Close()
		return
	}
}

func main() {
	logLevel := flag.Int("logLevel", 0, "an integer value(0 to 4)")
	flag.Parse()

	SetLogLevel(Level(*logLevel)) // initLog

	Trace.Println("MAIN Started")
	loop()
	Trace.Println("MAIN Completed")
}

func loop() {
	Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		Info.Println("Counter value is:", i)
	}
	Warning.Println("The counter variable is not being used")
	Trace.Println("Loop completed")
}
