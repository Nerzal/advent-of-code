package game

import (
	"regexp"
	"strings"
)

var pathRegex = regexp.MustCompile(`(\w{3})\s=\s\((\w{3}),\s(\w{3})\)`)

func FollowPath(instructions string, path map[string][]string) []int {

	var currentSigns []string

	for k, _ := range path {
		if strings.HasSuffix(k, "A") {
			currentSigns = append(currentSigns, k)
		}
	}

	var result []int

	for _, currentSign := range currentSigns {
		i := 0
		steps := 0

		for {
			steps++
			direction := instructions[i]

			if direction == 'R' {
				currentSign = path[currentSign][1]
			} else {
				currentSign = path[currentSign][0]
			}

			if strings.HasSuffix(currentSign, "Z") {
				result = append(result, steps)
				break
			}

			i++
			if i >= len(instructions) {
				i = 0
			}
		}
	}

	return result
}

// for {
// 	steps++
// 	direction := instructions[i]

// 	var nextSigns []string
// 	for _, currentSign := range currentSigns {
// 		if direction == 'R' {
// 			nextSigns = append(nextSigns, path[currentSign][1])
// 		} else {
// 			nextSigns = append(nextSigns, path[currentSign][0])
// 		}
// 	}

// 	signCounter := 0
// 	for _, sign := range nextSigns {
// 		if strings.HasSuffix(sign, "Z") {
// 			// println("reached Z after:", steps)
// 			signCounter++
// 		}
// 	}

// 	if signCounter != 0 && signCounter == len(nextSigns) {
// 		return steps
// 	}

// 	i++
// 	if i >= len(instructions) {
// 		i = 0
// 	}

// 	currentSigns = nextSigns
// }
// }

func GetPath(input []string) map[string][]string {
	result := map[string][]string{}

	for _, line := range input {
		destination, pathLeft, pathRight := LineToRoute(line)
		result[destination] = append(result[destination], pathLeft, pathRight)
	}

	return result
}

func LineToRoute(line string) (destination, pathLeft, pathRight string) {
	matches := pathRegex.FindStringSubmatch(line)

	destination = matches[1]
	pathLeft = matches[2]
	pathRight = matches[3]

	return
}
