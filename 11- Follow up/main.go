package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hellow, World!!")
	fmt.Println("Go" + "Lang")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("7.0 / 3.0 = ", fmt.Sprintf("%.2f",  7.0 / 3.0))
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(false && false)
}