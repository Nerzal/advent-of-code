package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle4/file"
	"github.com/Nerzal/advent-of-code/2023/puzzle4/part1/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	service := game.NewService(input)
	cards := service.ConvertToCards()

	result := game.Solve(cards)
	println(result)
}
