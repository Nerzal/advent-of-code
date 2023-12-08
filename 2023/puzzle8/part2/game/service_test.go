package game_test

import (
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle8/part2/game"
	"github.com/stretchr/testify/require"
)

func TestInput2(t *testing.T) {
	input := []string{
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	// gameService := game.NewService()
	path := game.GetPath(input)
	steps := game.FollowPath("LR", path)
	require.Equal(t, 6, steps)
}
