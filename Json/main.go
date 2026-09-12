package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"personName"`
	Age  int   
	City string
}

func main() {
	// person := Person{
	// 	Name: "john",
	// 	Age:  45,
	// 	City: "Dhaka",
	// }

	// rowJson, err := json.Marshal(person)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(rowJson))

	var p2 Person 
	jsontext := `{"personName":"john","Age":45,"City":"Dhaka"}`

	json.Unmarshal([]byte(jsontext),&p2)

	fmt.Println(p2)

}
