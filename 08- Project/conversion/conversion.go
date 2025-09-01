package conversion
import (
	"fmt"
	"strconv"
)
func StringsToFloat(lines []string)([]float64, error){
	
	// Convert lines into float64, then save it into slice
	prices := make([]float64, len(lines))
	for index, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error converting string to float:", err)
			return nil, err
		}
		prices[index] = price
	}
	return prices, nil
}