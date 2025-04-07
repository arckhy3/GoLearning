package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) LoadData() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New(err.Error())
	}
	// file.Close()

	return lines, nil
}

func (fm FileManager) WriteData(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		file.Close()
		return errors.New(err.Error())
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New(err.Error())
	}
	// file.Close()
	return nil
}
