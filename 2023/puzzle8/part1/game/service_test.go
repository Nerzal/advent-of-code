package game_test

import (
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle8/part1/game"
	"github.com/stretchr/testify/require"
)

func TestInput2(t *testing.T) {
	input := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	// gameService := game.NewService()
	path := game.GetPath(input)
	steps := game.FollowPath("RL", path)
	require.Equal(t, 2, steps)
}

func TestInput3(t *testing.T) {
	input := []string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	// gameService := game.NewService()
	path := game.GetPath(input)
	steps := game.FollowPath("LLR", path)
	require.Equal(t, 6, steps)
}
