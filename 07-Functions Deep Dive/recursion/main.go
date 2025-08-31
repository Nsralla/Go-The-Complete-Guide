package recursion
import "fmt"
func main(){
	fmt.Println(factorial(5))

	number := 7
	result := 1
	for i := 2; i<= number;i++{
		result *= i
	}
	fmt.Println(result)
}

func factorial(num int)int{
	if num == 1 || num == 0 {
		return 1
	}else if num < 0 {
		return -1
	}
	return num * factorial(num - 1) 
}

