package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Nerzal/advent-of-code/2023/puzzle1/part2/sum"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := sum.BuildSum(scanner)

	println(result)
}
