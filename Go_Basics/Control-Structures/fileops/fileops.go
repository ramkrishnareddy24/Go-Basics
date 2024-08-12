package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)



func WriteFloatToFile(value float64,fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	// Handling errors instead try{} catch
	if err != nil {
		return 1000, errors.New("failed to find value")
	}

	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return balance, nil
}