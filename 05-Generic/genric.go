package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, World!")

	// Example 1: Using the generic function with integers
	// T is inferred as 'int' because we pass integer literals 3 and 4
	value := add(3, 4)
	fmt.Println("Returned Value:", value)

	// Example 2: Using the generic function with float64
	// T is inferred as 'float64' because we pass float64 literals
	floatResult := add(3.14, 2.86)
	fmt.Println("Float Result:", floatResult)

	// Example 3: Using the generic function with strings
	// T is inferred as 'string' because we pass string literals
	// For strings, + means concatenation
	stringResult := add("Hello, ", "World!")
	fmt.Println("String Result:", stringResult)

	// Example 4: Explicit type specification (optional but sometimes useful)
	// We can explicitly tell Go what type T should be
	explicitResult := add[int](10, 20)
	fmt.Println("Explicit int Result:", explicitResult)

	// NOTE: This would cause a compile-time error because you can't mix types:
	// mixedResult := add(3, 3.14) // Error: int and float64 are different types
	// With generics, both parameters must be the same type T
}

// Generic function definition with type parameter T
// T is a type parameter that can be one of: float64, int, or string
// The syntax [ T float64 | int | string] defines a type constraint
// - T: is the name of the type parameter (can be any name, T is convention)
// - float64 | int | string: is a union constraint (T can be any of these types)
// - The | symbol means "OR" - T can be float64 OR int OR string
func add[T float64 | int | string](x, y T) T {
	// Both parameters x and y must be of the same type T
	// The return type is also T (same as the input parameters)
	// This ensures type safety - you can't mix different types
	return x + y

	// COMMENTED OUT: The old way of doing this without generics
	// This is what we would have had to write before Go 1.18 introduced generics
	// Notice how much longer and more error-prone this approach is:

	// We would need to use interface{} (or 'any') and type assertions
	// Type assertions check if a value is of a specific type at runtime
	// aInt, aIsInt := x.(int)
	// bInt, bIsInt := y.(int)
	// if aIsInt && bIsInt {
	// 	fmt.Println("Adding integers:", aInt, bInt)
	// 	// Perform addition
	// 	sum := aInt + bInt
	// 	fmt.Println("Sum:", sum)
	// 	return sum
	// }

	// aFloat, aIsFloat := x.(float64)
	// bFloat, bIsFloat := y.(float64)
	// if aIsFloat && bIsFloat {
	// 	fmt.Println("Adding floats:", aFloat, bFloat)
	// 	// Perform addition
	// 	sum := aFloat + bFloat
	// 	fmt.Println("Sum:", sum)
	// 	return sum
	// }

	// aString, aIsString := x.(string)
	// bString, bIsString := y.(string)
	// if aIsString && bIsString {
	// 	fmt.Println("Adding strings:", aString, bString)
	// 	// Perform addition
	// 	sum := aString + bString
	// 	fmt.Println("Sum:", sum)
	// 	return sum
	// }

	// return 0
	// BUT TOOOOO LONG!!!!
	// ^^ This demonstrates why generics are so powerful:
	// - No runtime type checking needed (compile-time safety)
	// - No repetitive code for each type
	// - Better performance (no type assertions at runtime)
	// - Clear, readable code that's easy to maintain

}
