package main
import "fmt"

func main(){
	age := 32
	// var agePointer *int // means that agePointer is pointing to an int
	// agePointer = &age
	// or a shortcut
	agePointer := &age

	fmt.Println("address where agePointer points: ", agePointer)
	fmt.Println("value at that address: ", *agePointer)
	fmt.Println("Age minus 18: ", calcAge((agePointer)))
}

func calcAge(age *int) int{
	return *age - 18
}