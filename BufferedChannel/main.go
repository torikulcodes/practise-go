package main

import (
	"fmt"
	"time"
)

func main() {

	var ch = make(chan string, 3)

	var totalTime = time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "file uploaded..."
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "file uploaded2..."
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file uploaded3..."
	}()

	for range 3 {
		fmt.Println(<-ch)
	}
	fmt.Println(time.Since(totalTime))
}
