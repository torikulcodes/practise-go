package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()
	uploadFile()
	saveToDb()
	sendEmail()
	fmt.Println("time taken", time.Since(start))
}


func uploadFile() {
	fmt.Println("uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done ")
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
