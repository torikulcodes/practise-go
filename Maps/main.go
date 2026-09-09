package main

import "fmt"


func main(){
	 myMap := make(map[string]int)

	myMap["user1score"] = 5
	myMap["user2score"] = 50

	fmt.Println(myMap["user2score"])
}