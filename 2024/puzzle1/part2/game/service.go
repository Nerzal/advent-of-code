package game

import (
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
	left := make([]int, 0, len(s.input))
	right := make([]int, 0, len(s.input))
	for _, line := range s.input {
		splitStrings := strings.Split(line, " ")

		num1, _ := strconv.Atoi(splitStrings[0])
		num2, _ := strconv.Atoi(splitStrings[3])

		left = append(left, num1)
		right = append(right, num2)
	}

	result := 0
	count := 0
	for _, numLeft := range left {
		for _, numRight := range right {
			if numLeft == numRight {
				count++
			}
		}
		result += numLeft * count
		count = 0
	}

	return result
}
