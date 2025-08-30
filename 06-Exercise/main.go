package main
import "fmt"

type Product struct {
	title string
	id int
	price float64 
}


func main(){
	fmt.Println("First task:")
	hobbies := []string{"Running", "Swimming", "Reading"}
	fmt.Println(hobbies)

	fmt.Println("\nSecond task:")
	fmt.Println("first element: ", hobbies[0])
	slicedHobbies := hobbies[1:3]
	fmt.Println("second and third combined: ", slicedHobbies)

	fmt.Print("\nThird task:")
	mainHobbies := hobbies[:2]
	fmt.Println("slice based on the first element that contains the first and second elements: ", mainHobbies)

	fmt.Print("\nFourth task:")
	hobbies_new := mainHobbies[1:3]
	fmt.Println("slice based on the second and last elements of the original array: ", hobbies_new)

	fmt.Print("\n Fifth task:")
	goals := []string{"Learn Go", "Build a web app"}
	goals[1] = "Build a REST API"
	goals = append(goals, "Contribute to open source")
	fmt.Println(goals)



	fmt.Print("\n Sixth task:")
	products := []Product{
		{title: "Phone", id: 0, price: 699.99},
		{title: "Tablet", id: 1, price: 399.99},
	}
	products = append(products, Product{title: "Laptop", id: 2, price: 999.99})
	fmt.Println(products)
}

// 1 - create a new array that contains 3 hobbies, then print it
// 2 - output the first element, second and third combined in one list
// 3 - create a slice based on the first element that contains the first and the second elmeents
// 4 - re slice the slice from 3 and change to contain the second and last element of the original array