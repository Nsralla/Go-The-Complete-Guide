package main
import "fmt"
// varidaic functions
func main() {
	// method 1
	total :=sumUp(1,21,23,55,100)
	fmt.Println("sum = ", total)

	//method2
	numbers := []int{1,2,3,4,5,6}
	total2 := sumUp(numbers...)
	fmt.Println("sum2 = ", total2)
}

func sumUp (numbers ...int)int {
	total := 0
	for _,num := range numbers {
		total += num
	}
	return total
}