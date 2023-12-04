package game_test

import (
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle3/part1/game"
)

func TestSolve(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	service := game.NewService(input)
	gameData := service.Convert()

	result := gameData.Solve()
	println(result)

	if result != 4361 {
		t.Error("result is wrong")
	}
}
