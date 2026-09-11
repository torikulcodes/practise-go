package main

import "fmt"

func pracInterface() {
	var data interface{}

	data = "mizan"
	fmt.Println(data)
	data = 25

	fmt.Println(data)
}

func emty(data any) {
	fmt.Println(data)
}

// type assertion

func typeAssertion(data any) {
	strData, ok := data.(string)

	if ok {
		fmt.Println(strData)

	}else{
		fmt.Println("This is not string")
	}

}

func main() {
	emty("torikul")
	pracInterface()
	typeAssertion("habib")
	typeAssertion(66)
}
