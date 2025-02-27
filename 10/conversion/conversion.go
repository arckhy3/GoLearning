package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(dataString []string) ([]float64, error) {
	var floatValues []float64

	for _, line := range dataString {
		floatValue, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New(err.Error())
		}
		floatValues = append(floatValues, floatValue)
	}

	return floatValues, nil
}
