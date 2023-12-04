package game_test

import (
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle2/part1/game"
	"github.com/stretchr/testify/require"
)

func TestConvertInput(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	converter := game.NewInputConverter(input)
	games := converter.Convert()
	require.NotNil(t, games)
	require.Equal(t, 1, games[0].ID)

	sum := 0

	for _, game := range games {
		if !game.Valid {
			continue
		}
		sum += game.ID
	}

	require.Equal(t, 8, sum)
}
