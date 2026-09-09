package main

import "fmt"

type Animal interface{
	speak()
}

type Dog struct {
}

type Cat struct {
}


func(d Dog) speak(){
	fmt.Println("woof woof")
}

func makeSound(d Animal){
	d.speak()
}

func makeSound2 (c Animal){

}

func main() {
	dexter := Dog{}

	makeSound(&dexter)
}
