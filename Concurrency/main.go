package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var fileUrl string

func main() {
	var start = time.Now()

	wg.Go(uploadFile)
	wg.Go(saveToDb)
	wg.Go(sendEmail)
	wg.Wait()

	fmt.Println("url",fileUrl)

	fmt.Println("time taken", time.Since(start))
}

func uploadFile()  {
	fmt.Println("uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done ")

	 fileUrl  = "https///pro.com"


}

func saveToDb() {
	fmt.Println("save db file...")
	time.Sleep(1 * time.Second)
	fmt.Println("File save done ")

}

func sendEmail() {
	fmt.Println("sending email")
	time.Sleep(2 * time.Second)
	fmt.Println("sending email done ")
}
