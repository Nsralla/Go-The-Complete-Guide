package anonyms

import "fmt"

// Anonymous functions
func main() {
	func() {
		println("Hello from an anonymous function")
	}()
	numbers := []int{1,2,3,4,5}
	new_numbers := transformNumbers(&numbers, func (value int) int {
		return value * 2
	})
	fmt.Println(*new_numbers)
}


func transformNumbers(numbers *[]int, transformFunc func(int) int) *[]int{
	transformedNumbers := []int{}

	for _, value := range *numbers{
		transformedNumbers = append(transformedNumbers, transformFunc(value))
	}
	return &transformedNumbers
}