package game

import (
	"math"
	"strconv"
	"strings"
)

type Service struct {
	input []string
}

func NewService(input []string) *Service {
	return &Service{
		input: input,
	}
}

func (s *Service) Run() int {
	var result int

	for _, line := range s.input {
		lineNumbers := toIntArray(line)

		if isSafe(lineNumbers) {
			result++
			continue
		}

		for i := 0; i < len(lineNumbers); i++ {
			modifiedList := append([]int{}, lineNumbers[:i]...)
			modifiedList = append(modifiedList, lineNumbers[i+1:]...)

			if isSafe(modifiedList) {
				result++
				break
			}
		}
	}

	return result
}

func isSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(numbers); i++ {
		diff := math.Abs(float64(numbers[i]) - float64(numbers[i-1]))

		if diff < 1 || diff > 3 {
			return false
		}

		if numbers[i] > numbers[i-1] {
			decreasing = false
		}

		if numbers[i] < numbers[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func toIntArray(line string) []int {
	splitLines := strings.Split(line, " ")
	var result []int

	for _, numStr := range splitLines {
		if numStr == "" {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		result = append(result, num)
	}

	return result
}
