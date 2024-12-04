package main

import (
	"github.com/Nerzal/advent-of-code/2024/puzzle1/file"
	"github.com/Nerzal/advent-of-code/2024/puzzle1/part2/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	converter := game.NewService(input)
	result := converter.Run()

	println(result)
}
