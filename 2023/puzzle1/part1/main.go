package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line)
	}

	println(sum)
}

func handleLine(line string) int {
	var (
		firstDigit       rune
		lastDigit        rune
		foundSecondDigit bool
	)

	for _, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}

		if firstDigit == 0 {
			firstDigit = c
			continue
		}

		lastDigit = c
		foundSecondDigit = true
	}

	if !foundSecondDigit {
		lastDigit = firstDigit
	}

	bothDigits, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(firstDigit), string(lastDigit)))
	return bothDigits
}
