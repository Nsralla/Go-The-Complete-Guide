package funs
import "fmt"

func main() {
	// This is the entry point of the program
	numbers := []int{1, 2, 3, 4}
	// fmt.Println(*dNumbers(&numbers))
	fmt.Println(*transformNumbers(&numbers, double))
	fmt.Println(*transformNumbers(&numbers, tripple))
}

// func dNumbers(numbers *[]int) *[]int {
// 	dnumbers := []int{}
// 	for _, num := range *numbers {
// 		dnumbers = append(dnumbers, num*2)
// 	}
// 	return &dnumbers
// } 


func transformNumbers(numbers *[]int, transform func(int) int) *[]int{
	transformedNumbers := []int{}
	for _, num := range *numbers{
		transformedNumbers = append(transformedNumbers, transform(num))
	}
	return &transformedNumbers
}
func double(value int) int {
	return value * 2
}

func tripple(value int) int {
	return value * 3
}