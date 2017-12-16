package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"reflect"
)

func ReadFile(filename string) ([]byte , error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main () {
	fname := "/Users/nipun.jindal/Documents/testing_go_lang_file_io"
	data, _ := ReadFile(fname)
	fmt.Println(len(data))
	fmt.Println("datatype of ioutil.ReadAll : ", reflect.TypeOf(data))
	fmt.Println(data)  // will print the byte values a int
	fmt.Printf("As string : %s", data)

}
