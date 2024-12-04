package game

import (
	"regexp"
	"strconv"
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
	result := 0
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, line := range s.input {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			m1, _ := strconv.Atoi(match[1])
			m2, _ := strconv.Atoi(match[2])
			result += m1 * m2
		}
	}

	return result
}
