package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func ReadFile(file_path string) ([]byte, error){
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func main() {
	file_path := "/Users/nipun.jindal/Documents/testing_go_lang_file__io" // wrong file_path will throw error
	data, _ := ReadFile(file_path)
	fmt.Println(data) // I will never come to this `coz panic stopped execution of Program
}
