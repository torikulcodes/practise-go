package main

import (
	"fmt"
)

func process(callback func()) {
	callback()
}

func main() {

	sayhello := func() {

		fmt.Println("hello there")
	}
	process(sayhello)
}

//  anonymous callback function

