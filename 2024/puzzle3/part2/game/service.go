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

	regexMul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	regexDo := regexp.MustCompile(`do\(\)`)
	regexDont := regexp.MustCompile(`don't\(\)`)

	mulEnabled := true

	// FML

	for _, line := range s.input {
		for i := 0; i < len(line); i++ {
			if regexMul.MatchString(line[i:]) && mulEnabled {
				matches := regexMul.FindStringSubmatch(line[i:])
				m1, _ := strconv.Atoi(matches[1])
				m2, _ := strconv.Atoi(matches[2])
				result += m1 * m2
			}

			if regexDo.MatchString(line[i:]) {
				mulEnabled = true
			} else if regexDont.MatchString(line[i:]) {
				mulEnabled = false
			}
		}
	}

	return result
}
