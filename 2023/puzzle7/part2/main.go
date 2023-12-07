package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle7/part2/game"
	"github.com/Nerzal/advent-of-code/2023/template/file"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	gameService := game.NewService(input)
	hands := gameService.ToHands()

	game.SortHands(hands)

	result := game.Play(hands)
	println(result)

	// 250035606 <- too low
}
