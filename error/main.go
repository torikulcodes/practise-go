package main

import (
	"errors"
	"fmt"
)

func main() {
	result, error := divide(10, 0)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
}

func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
