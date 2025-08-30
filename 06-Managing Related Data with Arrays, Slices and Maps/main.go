package main

import "fmt"

func main(){
	// Building Dynamic Lists
	prices := []float64{10.99, 3.54}
	fmt.Print("prices[1]: ", prices[1])
	// fmt.Println("outof index: ", prices[2]) // This will cause a runtime error: index out of range
	prices = append(prices, 4.99) // it doesn't edit the original array, it creates a new one
	// adding multiple elements
	prices = append(prices, 5.99, 6.99, 7.99)
	fmt.Println("\nAfter appending, prices: ", prices)
	// what happend to original prices? GO will garbage collect the original slice if there are no references to it

}

// func main() {
// 	prices := []float64{1.0, 2.0, 3.0, 3.99, 4.99}
// 	var productNames = [4]string{"MacBook2"}
// 	productNames[0] = "MacBook3"
// 	productNames[1] = "Iphone"
// 	productNames[2] = "iPad"
// 	// productNames[3] = "Apple Watch"

//     // SLICING
// 	// Slicing doesn't make copy, it creates a reference to the original array
// 	fmt.Println("SLICING ******")
// 	featuredPrices := prices[1:3]
// 	fmt.Println("featuredPrices: ", featuredPrices)
// 	fmt.Println("prices[:3]: ", prices[:3])
// 	// If you modify the original slice, the changes will be reflected in the sliced version
// 	prices[1] = 2.5
// 	fmt.Println("Prices array has been modified, first element became: 2.5 -> featured prices :")
// 	fmt.Println("featuredPrices: ", featuredPrices)
// 	// Also changing featuredPrices will reflect on prices
// 	featuredPrices[0] = 2.8
// 	fmt.Println("Prices array has been modified, first element became: 2.8 -> featured prices :")
// 	fmt.Println("featuredPrices: ", featuredPrices)
// 	fmt.Println("prices: ", prices)

// 	fmt.Println("Array and Slice Basics ******")
// 	fmt.Println("prices[1:]: ", prices[1:])
// 	fmt.Println("prices[:]: ", prices[:])
// 	// No Negative indexing
// 	// fmt.Println(prices[-1]) // This will cause a compile-time error

// 	fmt.Println("Printing ******")
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(productNames[3])

// 	fmt.Println("Length ************")
// 	fmt.Println("Length of prices:", len(prices)) // Length is the number of elements in the slice
// 	fmt.Println("Capacity of prices:", cap(prices)) // Capacity is the total allocated space
// 	fmt.Println("length of featuredPrices:", len(featuredPrices))
// 	fmt.Println("capacity of featuredPrices:", cap(featuredPrices))

// 	// for i, price := range prices {
// 	// 	fmt.Printf("Item %d: $%.2f\n", i+1, price)
// 	// }
// }