package main

import (
	"encoding/json"
	"fmt"
)

//Employee Struct
type Employee struct {
	ID int
	FirstName, LastName, Job string
}

func main () {
	emp := Employee{ID : 23, FirstName:"Nipun", LastName:"Jindal", Job:"Devloper"}
	//Encoding Json
	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr := string(data)
	fmt.Println(jsonStr) // the json string
	// Decode json byte-data
	b := []byte(`{"ID":23,"FirstName":"Nipun","LastName":"Jindal","Job":"Devloper"}`) // we are using "apostrophe" for covering the string`s double-Quotes
	var emp2 Employee
	// Decoding json data to thte "emp2" Employee
	err2 := json.Unmarshal(b, &emp2) // the value is decoded to "emp2"
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Printf("Employee (ID:%d, FirstName:%s, LastName:%s, Job:%s)", emp2.ID, emp2.FirstName, emp2.LastName, emp2.Job)
}
