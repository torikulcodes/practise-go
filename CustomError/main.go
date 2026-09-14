package main

import "fmt"

type CustomError struct{
	message string
	code int
}

func (cu *CustomError)Error() string{
	return cu.message
}

func login(password string)error{
 if(password != "1234"){
	return &CustomError{
		message: "password do not match",
		code: 401,
	}
 }
 return nil
}



func main(){
 var err = login("2345")
 if(err != nil){
  fmt.Println("error",err,"code",err.(*CustomError).code)
 }
 fmt.Println("main ends")
}

