package sum

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func BuildSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += HandleLine(line)
	}

	return sum
}

func HandleLine(line string) int {
	digits := getDigits(line)

	if len(digits) == 0 {
		return 0
	}

	// sort digits by Index ascending
	slices.SortFunc(digits, func(a, b digit) int {
		switch {
		case a.Index < b.Index:
			return -1
		case a.Index > b.Index:
			return 1
		default:
			return 0
		}
	})

	digitString := fmt.Sprintf("%d%d", digits[0].Value, digits[len(digits)-1].Value)
	digit, _ := strconv.Atoi(digitString)

	// println("Result: ", digit, "String:", line)

	return digit
}

func getDigits(line string) []digit {
	return append(checkDigits(line), checkWords(line)...)
}

func checkDigits(line string) []digit {
	var digits []digit

	for i, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}

		intValue, _ := strconv.Atoi(string(c))

		digit := digit{
			Index: i,
			Value: intValue,
		}

		digits = append(digits, digit)
	}

	return digits
}

var numberWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func checkWords(line string) []digit {
	var result []digit

	for i, word := range numberWords {
		index := strings.Index(line, word)
		if index == -1 {
			continue
		}

		result = append(result, digit{Index: index, Value: i})

		lastIndex := strings.LastIndex(line, word)
		if index == lastIndex {
			continue
		}

		result = append(result, digit{Index: lastIndex, Value: i})

	}

	return result
}
