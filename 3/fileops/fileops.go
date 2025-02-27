package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFLoatToFile(fileName string, balance float64) {
	valueText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
func ReadFloatFromFile(fileName string) (value float64, errMsg error) {
	value = 1000
	errMsg = nil
	data, err := os.ReadFile(fileName)
	if err != nil {
		errMsg = errors.New("Failed to find the file.")
		return
	}
	valueText := string(data)
	value, err = strconv.ParseFloat(valueText, 64)
	if err != nil {
		errMsg = errors.New("Failed to convert data to float.")
		return
	}
	return
}
