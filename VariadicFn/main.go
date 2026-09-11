package main

import "fmt"

// Must be the last parameter: If your function accepts regular parameters alongside a variadic one, the variadic parameter must be placed at the very end of the function signature.

func sum (nums ...int) int{
  total := 0

  for _, num:= range nums {
	 total +=num
  }

  return  total
}


func main(){

	 fmt.Println(sum(56,77,55))
	 fmt.Println(sum(8))
}
