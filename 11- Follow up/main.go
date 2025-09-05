package main

import (
	"fmt"

)


func main(){
	var a [5]int
	fmt.Println(a)
	for i := range 5{
		a[i] = i + 1
	}
	fmt.Println(a)
	fmt.Println(len(a))
	// Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.

	var arr = [7]int{1,11,111,1111} // rest will be zero
	fmt.Println(arr)
	fmt.Println(len(arr))

	var arr2 = [...]int{100,1000,2000}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	b := [...]int{100, 3: 400, 500, 150}
	fmt.Println("idx:", b)

	var twoD = [2][2]int{
		{6,11},
		{87,32},
	}

	fmt.Println(twoD)

	// var x [3]int
	// x =[...]int{1,2,3}

	var twoD2 [2][3]int
	twoD2 = [...][3]int{
		{10,100,1000},
		{20,200,2000},
	}
	fmt.Println(twoD2)

}