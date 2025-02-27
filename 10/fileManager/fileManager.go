package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func LoadData(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New(err.Error())
	}
	file.Close()

	return lines, nil
}

func WriteData(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		file.Close()
		return errors.New(err.Error())
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New(err.Error())
	}
	file.Close()
	return nil
}
