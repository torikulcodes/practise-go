package main

import "fmt"

func main() {
	// myMap := map[string]string{
	// 	"name":    "torikul",
	// 	"success": "ok",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }

	// myarr := []string{
	// 	"green",
	// 	"red",
	// 	"yellow",
	// }

	// for key, value := range myarr {
	// 	fmt.Println(key, value)
	// }

name := "next level"

for key, value := range name {
    fmt.Printf("%d %c\n", key, value)
}
}
