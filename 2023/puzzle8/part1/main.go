package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle8/file"
	"github.com/Nerzal/advent-of-code/2023/puzzle8/part1/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	path := game.GetPath(input)
	steps := game.FollowPath("LRRRLRRRLRRLRRLRLRRLRRLRRRLRRLRRRLRRRLLRRRLRRRLRRRLRLRRLRRRLRLRRRLRRRLLRLRLRRLRRLLLRRLRRLRRRLLRRRLLRRRLRLRRRLRRRLLRRLRLLRLRRRLRRLRRLRLRLRLRLRLRRRLRLRRRLLRLRRLRRRLRRRLRLRRLRLLLRLRLRLRLRLRRRLLRRLRLRLLRRRLRRLRRRLRRLRRLRRRLLRRLRLRRLRRRLRRLRLRRLRLLRRLLRLRRRLRRLRLLRRRR", path)
	println(steps)
}
