package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID int `json:"id,omitempty"` // this will tag ID `s value with id tag and omit them if it is empty, best part : don`t put a space b/w tag and omitempty eg : `json:"id, omitempty"` won`t work
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Job string `json:"job"`
}

type Employee2 struct {
	ID int `json:"id"` // this will tag ID `s value with id tag and omit them if it is empty, best part : don`t put a space b/w tag and omitempty eg : `json:"id, omitempty"` won`t work
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Job string `json:"job"`
}

//Experiment if I hav a struct with only some field only with json tags, not all

func main() {
	emp1 := Employee{
		ID:12,
		FirstName:"Nipun",
		LastName:"Jindal",
		Job:"Developer",
	}
	jsonStr, err := json.Marshal(emp1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jsonStr)) // ---------1

	//checking "omitempty" in json
	emp2 := Employee{
		FirstName:"Nipun",
		LastName:"Jindal",
		Job:"developer",
	}
	jsonStr, err = json.Marshal(emp2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jsonStr)) // ----------2

	err = json.Unmarshal(jsonStr, &emp2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(emp2) // ---------------------3
	// checking for what happen if I don`t have omitEmpty, I think it will put some default value for that data-type
	emp3 := Employee2{
		FirstName:"Nipun",
		LastName:"Jindal",
		Job:"developer",
	}
	data, err := json.Marshal(emp3) // still it will compile, even though err is declared before
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(data)) // ----------------- 4
}
