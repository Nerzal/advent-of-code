package main

import (
	"fmt"

	"github.com/Nerzal/advent-of-code/2023/puzzle2/part1/game"
	"github.com/Nerzal/advent-of-code/2023/template/file"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	converter := game.NewInputConverter(input)
	fmt.Println(converter)
}
