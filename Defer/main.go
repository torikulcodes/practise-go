package main

import "fmt"

// 💡 Core Rules of DeferTo use defer effectively, you must understand three foundational behaviors:1. Last-In, First-Out (LIFO) OrderIf you use multiple defer statements in a single function, Go pushes them onto a stack. When the surrounding function exits, the deferred calls execute in reverse order.

func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")

    fmt.Println("Main function body")
}

// Output:
// Main function body
// Third defer
// Second defer
// First defer


// 2. Arguments are Evaluated ImmediatelyWhen a defer statement is evaluated, any arguments passed to the deferred function are calculated right then, not when the function actually executes.

func argEvaluated() {
    x := 10
    defer fmt.Println("Value of x in defer:", x) // captures x = 10

    x = 20
    fmt.Println("Value of x in main:", x)
}

// Output:
// Value of x in main: 20
// Value of x in defer: 10


// 3. Defer Can Modify Named Return ValuesIf the surrounding function uses named return values, a deferred anonymous function can access and modify those values before they are sent back to the caller.

func tripleValue() (result int) {
    defer func() {
        result += 5 // modifies the named return variable
    }()
    return 10 // result is initially set to 10
}

// Calling tripleValue() returns 15
