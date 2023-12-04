package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle3/file"
	"github.com/Nerzal/advent-of-code/2023/puzzle3/part1/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	service := game.NewService(input)
	gameData := service.Convert()

	result := gameData.Solve()
	println(result)
}
