package closure

import "fmt"

// closure
func main() {
	func() {
		println("Hello from an anonymous function")
	}()
	numbers := []int{1,2,3,4,5}

	doubleFunc := createTransformFunc(2)
	tripple := createTransformFunc(3)
	
	new_numbers := transformNumbers(&numbers, doubleFunc)
	new_numbers2 := transformNumbers(&numbers, tripple)
	fmt.Println(*new_numbers)
	fmt.Println(*new_numbers2)
}


func transformNumbers(numbers *[]int, transformFunc func(int) int) *[]int{
	transformedNumbers := []int{}

	for _, value := range *numbers{
		transformedNumbers = append(transformedNumbers, transformFunc(value))
	}
	return &transformedNumbers
}

func createTransformFunc(factor int) func(int)int{
	return func (value int) int{
		return value * factor
	}
}