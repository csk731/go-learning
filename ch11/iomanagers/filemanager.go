package iomanagers

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	// fmt.Println("Reading file", fm.InputFilePath)
	file, err := os.Open(fm.InputFilePath)
	// fmt.Println(err)
	if err != nil {
		return nil, errors.New("failed to open a file, try again")
	}
	scanner := bufio.NewScanner(file)
	var lines []string = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, errors.New("failed to read the file, try again")
	}
	file.Close()
	return lines, nil
}

func (fm *FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create a file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	file.Close()
	return nil
}

func NewFileManager(inputPath string, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
