package main

import (
	"encoding/json"
	"fmt"
)

type LoginInfo struct {
	User string `json:"user"`
	Password string `json:"-"`
}

func main () {
	u := LoginInfo {
		User:"Nipun",
		Password:"temp",
	}
	jsonStr, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jsonStr)) // As we see, password field has been not included in the json
}
