package file

import (
	"bufio"
	"os"
)

func GetInputData() ([]string, error) {
	file, err := os.Open("input")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
