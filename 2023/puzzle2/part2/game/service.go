package game

import (
	"regexp"
	"strconv"
	"strings"
)

type InputConverter struct {
	input []string
}

func NewInputConverter(input []string) *InputConverter {
	return &InputConverter{
		input: input,
	}
}

var idRegex = regexp.MustCompile(`Game\s(?P<ID>\d+)`)

func (c *InputConverter) Convert() []*Game {
	result := make([]*Game, len(c.input))

	for i, line := range c.input {
		splitStrings := strings.Split(line, ":")

		idLine := splitStrings[0]
		id := getGameID(idLine)

		rounds := getRounds(splitStrings[1])

		g := &Game{
			ID:     id,
			Rounds: rounds,
		}
		g.Validate()

		result[i] = g
	}

	return result
}

func getRounds(line string) []*Round {
	var result []*Round

	rounds := strings.Split(line, ";")

	for _, roundLine := range rounds {
		round := getRound(roundLine)
		round.Validate()
		result = append(result, round)
	}

	return result
}

var roundRegex = regexp.MustCompile(`(\d+)`)

func getRound(line string) *Round {
	result := &Round{}

	numberStrings := strings.Split(line, ",")
	for _, numberString := range numberStrings {
		matches := roundRegex.FindStringSubmatch(numberString)
		number, _ := strconv.Atoi(matches[0])

		switch {
		case strings.Contains(numberString, "red"):
			result.Red = number
		case strings.Contains(numberString, "blue"):
			result.Blue = number
		case strings.Contains(numberString, "green"):
			result.Green = number
		}
	}

	return result
}

func getGameID(line string) int {
	matches := idRegex.FindStringSubmatch(line)
	id, _ := strconv.Atoi(matches[1])
	return id
}
