package game

import (
	"regexp"
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

func (s *Service) ConvertToCards() []Card {
	var result []Card

	for _, line := range s.input {
		card := convertToCard(line)
		result = append(result, card)
	}

	return result
}

var numberRegExp = regexp.MustCompile(`(\d+)`)

func convertToCard(line string) Card {
	splittedStrings := strings.Split(line, ":")

	cardIDLine := splittedStrings[0]
	id := getCardID(cardIDLine)

	numberStrings := strings.Split(splittedStrings[1], "|")
	winningNumbers := numberStringsToSlice(numberStrings[0])
	chosenNumbers := numberStringsToSlice(numberStrings[1])

	result := Card{
		ID:             id,
		Numbers:        chosenNumbers,
		WinningNumbers: winningNumbers,
	}

	return result
}

func getCardID(line string) int {
	matches := numberRegExp.FindStringSubmatch(line)
	id, _ := strconv.Atoi(matches[0])
	return id
}

func numberStringsToSlice(numberStrings string) []int {
	matches := numberRegExp.FindAllStringSubmatch(numberStrings, -1)

	var result []int

	for _, match := range matches {
		number, _ := strconv.Atoi(match[0])
		result = append(result, number)
	}

	return result
}
