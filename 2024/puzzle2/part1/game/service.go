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

		lastNum := lineNumbers[0]

		increasing := true
		decreasing := true
		shortDistance := true

		for i, num := range lineNumbers {
			if i == 0 {
				continue
			}

			if num == lastNum {
				increasing = false
				decreasing = false
			}

			if num > lastNum {
				decreasing = false
			}

			if num < lastNum {
				increasing = false
			}

			abs := math.Abs(float64(num) - float64(lastNum))
			if abs < 1 || abs > 3 {
				shortDistance = false
			}

			lastNum = num
		}

		if shortDistance {
			if increasing || decreasing {
				result++
			}
		}
	}

	return result
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
