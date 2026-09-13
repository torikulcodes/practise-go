package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

var counter int

func main() {

	var start = time.Now()

	for range 1000 {
		wg.Go(increment)
	}

	wg.Wait()
	fmt.Println("counter value is", counter)

	fmt.Println("time taken", time.Since(start))

}

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter = counter + 1
	fmt.Println(counter)
}
