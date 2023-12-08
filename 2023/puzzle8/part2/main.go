package main

import (
	"fmt"

	"github.com/Nerzal/advent-of-code/2023/puzzle8/file"
	"github.com/Nerzal/advent-of-code/2023/puzzle8/part2/game"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	path := game.GetPath(input)
	steps := game.FollowPath("LRRRLRRRLRRLRRLRLRRLRRLRRRLRRLRRRLRRRLLRRRLRRRLRRRLRLRRLRRRLRLRRRLRRRLLRLRLRRLRRLLLRRLRRLRRRLLRRRLLRRRLRLRRRLRRRLLRRLRLLRLRRRLRRLRRLRLRLRLRLRLRRRLRLRRRLLRLRRLRRRLRRRLRLRRLRLLLRLRLRLRLRLRRRLLRRLRLRLLRRRLRRLRRRLRRLRRLRRRLLRRLRLRRLRRRLRRLRLRRLRLLRRLLRLRRRLRRLRLLRRRR", path)

	fmt.Println(steps)

}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
