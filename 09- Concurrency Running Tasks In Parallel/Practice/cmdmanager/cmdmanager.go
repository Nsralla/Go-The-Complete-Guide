package cmdmanager

import "fmt"

type CMDManager struct {
}

// READ LINES FROM CMD
func (cmd *CMDManager) ReadLinesFromFile() (lines []string, err error) {
	for {
		var line string
		fmt.Print("Enter a price (or type 'done' to finish): ")
		_, err := fmt.Scanln(&line)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if line == "done" {
			break
		}
		lines = append(lines, line)
	}

	return lines, nil
}

// WRITE DATA TO CMD	
func (cmd *CMDManager) WriteJson(data any) error {
	fmt.Println("Writing JSON to file...")
	fmt.Println(data)
	fmt.Println("JSON written successfully.***************")
	return nil
}


func New() *CMDManager {
	return &CMDManager{}
}