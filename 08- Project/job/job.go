package job

import (
	"errors"
	"fmt"
	"math"

	"example.com/calc/conversion"
	// "example.com/calc/filemanager"
	"example.com/calc/iom"
)

type Job struct {
	TaxRate            float64            `json:"tax_rate"`        // tax rate
	Prices             []float64          `json:"prices"`          // array containing prices for the job
	PricesIncludingTax map[string]float64 `json:"prices_with_tax"` // dictionary for each price with its tax included (price * tax). key: price, value: price * tax
	// JSON object keys are always strings.
	IOManager iom.IOManager `json:"-"` // this tells go how we want job to appear in the json file
}

// Compute the PricesIncludingTax
func (job *Job) Process() error{

	// Read prices from file
	err := job.loadPrices()
	if err != nil {
		// fmt.Println("Error loading prices:", err)
		return errors.New("failed to load prices " + err.Error())
	}

	fmt.Print("Prices loaded successfully \n *********************\n")

	// Compute prices including tax
	result := make(map[string]float64) // Dictionary to temporarily hold `prices including tax`
	for _, price := range job.Prices {
		priceWithTax := (1 + job.TaxRate) * price
		result[fmt.Sprintf("%.2f", price)] = math.Round(priceWithTax*100) / 100 // to round to 2 decimal places
	}
	// Assign the `Prices including tax` to the Job struct
	job.PricesIncludingTax = result
	// Write `price with tax` to json file
	job.IOManager.WriteJson(job)

	return nil

}

// Read prices from .txt file, Then assign it to the job struct
func (j *Job) loadPrices() error {

	// Read prices from prices.txt
	// Reading logic exists at example.com/calc/filemanager
	lines, err := j.IOManager.ReadLinesFromFile()
	if err != nil {
		// fmt.Println("Error reading lines from file:", err)
		return errors.New("failed to read lines from file: " + err.Error())
	}

	// Convert lines into float64, then save it into slice.
	// Conversion functions is at example.com/calc/conversion. will return prices, or an error if exists
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("Error converting strings to float:", err)
		return errors.New("failed to convert strings to float: " + err.Error())
	}

	// Assign the prices for the Job struct
	j.Prices = prices
	return nil
}

// Constructor to create a new Job with IOManager
func New(taxRate float64, ioM iom.IOManager) *Job {
	return &Job{
		TaxRate:   taxRate,
		IOManager: ioM,
	}
}
