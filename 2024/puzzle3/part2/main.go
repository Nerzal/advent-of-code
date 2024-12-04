package main

import (
	"github.com/Nerzal/advent-of-code/2024/puzzle3/file"
	"github.com/Nerzal/advent-of-code/2024/puzzle3/part2/game"
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
