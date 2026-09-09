package main

import "fmt"

type user struct {
	name       string
	age        int
	isLoggedIn bool
	greets      func()
}

func main() {

	// user1 := user{
	// 	name: "torikul",
	// 	age: 55,
	// 	isLoggedIn: false,
	// }
	// user1.greet = func() {
	// 	fmt.Println("hello",user1.name)
	// }

	// user1.greet()

	// receiver function

	user1 := user{
		name:       "torikul",
		age:        55,
		isLoggedIn: false,
	}

	user1.greet()

}
func (u user) greet() {
 fmt.Println("hello",u.name)
}
