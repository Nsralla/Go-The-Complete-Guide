package job

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"example.com/calc/conversion"
)

type Job struct {
	TaxRate            float64 // tax rate
	Prices             []float64 // array containing prices for the job
	PricesIncludingTax map[float64]float64 // dictionary for each price with its tax included (price * tax). key: price, value: price * tax
}

// Compute the PricesIncludingTax
func (job *Job) Process() { 

	// Read prices from file
	job.loadPrices()
	fmt.Println("Prices loaded from file: ", job.Prices)
	fmt.Print("Prices loaded successfully \n *********************\n")

	// Compute prices including tax
	result := make(map[float64]float64) // Dictionary to temporarily hold prices including tax
	for _, price := range job.Prices {
		priceWithTax := (1 + job.TaxRate) * price
		result[price] = math.Round(priceWithTax * 100) / 100 // to round to 2 decimal places
	}
	// Assign the Prices including tax to the Job struct
	job.PricesIncludingTax = result
	// Print it
	PrintPricesWithTax(job)

}

// Read prices from .txt file, Then assign it to the job struct
func (j *Job) loadPrices(){
	// Open the file
	file, err := os.Open("job/prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Read the lines from the file
	lines := []string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() { // this will read line by line
		line := reader.Text()
		lines = append(lines, line)
	}


	// Convert lines into float64, then save it into slice.
	// Conversion functions is at example.com/calc/conversion. will return prices, or an error if exists
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("Error converting strings to float:", err)
		file.Close()
		return
	}

	// Assign the prices for the Job struct
	j.Prices = prices
	
	
	err = reader.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}
	file.Close()
}

// Constructor to create a new Job
func New(taxRate float64) *Job {
	return &Job{
		TaxRate: taxRate,
		// Prices:  []float64{10, 20, 30},
	}
}

// Print the prices including tax
func PrintPricesWithTax(job *Job){
	fmt.Printf("Tax rate : %.2f has the following prices :  ", job.TaxRate)
	for _, value := range job.PricesIncludingTax {
		fmt.Printf("%.2f, ", value)
	}
	fmt.Println()
}