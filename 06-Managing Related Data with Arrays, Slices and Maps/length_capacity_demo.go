	package main

	import "fmt"

	func main33() {
		fmt.Println("=== LENGTH vs CAPACITY DEMONSTRATION ===\n")

		// 1. Creating a slice with make()
		fmt.Println("1. Using make() to create slice:")
		numbers := make([]int, 3, 5) // length=3, capacity=5
		fmt.Printf("numbers: %v\n", numbers)
		fmt.Printf("Length: %d, Capacity: %d\n\n", len(numbers), cap(numbers))

		// 2. Adding elements to see capacity in action
		fmt.Println("2. Adding elements:")
		numbers[0] = 10
		numbers[1] = 20
		numbers[2] = 30
		fmt.Printf("After filling: %v\n", numbers)
		fmt.Printf("Length: %d, Capacity: %d\n\n", len(numbers), cap(numbers))

		// 3. Using append() - length increases, capacity may increase
		fmt.Println("3. Using append():")
		numbers = append(numbers, 40)
		fmt.Printf("After append(40): %v\n", numbers)
		fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

		numbers = append(numbers, 50)
		fmt.Printf("After append(50): %v\n", numbers)
		fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

		// This will exceed capacity and trigger reallocation
		numbers = append(numbers, 60)
		fmt.Printf("After append(60): %v\n", numbers)
		fmt.Printf("Length: %d, Capacity: %d\n\n", len(numbers), cap(numbers))

		// 4. Slicing and its effect on capacity
		fmt.Println("4. Slicing effects:")
		original := []int{1, 2, 3, 4, 5, 6, 7, 8}
		fmt.Printf("Original: %v (len=%d, cap=%d)\n", original, len(original), cap(original))

		slice1 := original[2:5]
		fmt.Printf("slice1 [2:5]: %v (len=%d, cap=%d)\n", slice1, len(slice1), cap(slice1))

		slice2 := original[4:6]
		fmt.Printf("slice2 [4:6]: %v (len=%d, cap=%d)\n", slice2, len(slice2), cap(slice2))

		slice3 := original[0:3]
		fmt.Printf("slice3 [0:3]: %v (len=%d, cap=%d)\n\n", slice3, len(slice3), cap(slice3))

		// 5. What happens when you modify?
		fmt.Println("5. Modifying slices:")
		fmt.Println("Before modification:")
		fmt.Printf("original: %v\n", original)
		fmt.Printf("slice1: %v\n", slice1)

		slice1[0] = 999 // This changes original[2]
		fmt.Println("After slice1[0] = 999:")
		fmt.Printf("original: %v\n", original)
		fmt.Printf("slice1: %v\n", slice1)
	}
