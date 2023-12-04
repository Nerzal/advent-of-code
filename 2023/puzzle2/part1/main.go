package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle2/file"
	"github.com/Nerzal/advent-of-code/2023/puzzle2/part1/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	converter := game.NewInputConverter(input)
	games := converter.Convert()

	sum := 0

	for _, game := range games {
		if !game.Valid {
			continue
		}
		sum += game.ID
	}

	println(sum)
}
