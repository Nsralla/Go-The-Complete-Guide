package main

import (
	"fmt"
)


func main(){
	call_func := outer()
	fmt.Println(call_func())
	fmt.Println(call_func())
	fmt.Println(call_func())

	call2 := outer()
	fmt.Println(call2())
}


func outer() func () int{
	i := 0
	return func()(int){
		i++
		return i
	}
}