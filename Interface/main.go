package main

import "fmt"

// 1. Define the interface contract
type Greeter interface {
    Greet() string
}

// 2. Implement the method on a concrete type (Struct)
type Person struct {
    Name string
}

// Person implicitly implements Greeter because it has the Greet() method
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// 3. Consume the interface
func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    p := Person{Name: "Alice"}
    SayHello(p) // Works perfectly!
}
