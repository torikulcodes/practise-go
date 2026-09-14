package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)
	go uploadFile(ch)
	fileUrl := <-ch
	fmt.Println("file url", fileUrl)
}

func uploadFile(c chan string) {
	fmt.Println("file upload...")
	time.Sleep(3 * time.Second)

	fmt.Println("file uploading done")
	fileUrl := "https://s3.4545453"
	c <- fileUrl
}
