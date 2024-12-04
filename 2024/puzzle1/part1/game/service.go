package game

import (
	"math"
	"slices"
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

	slices.Sort(left)
	slices.Sort(right)

	var result int
	for i := 0; i < len(left); i++ {
		distance := math.Abs(float64(left[i]) - float64(right[i]))
		result += int(distance)
	}

	return result
}
