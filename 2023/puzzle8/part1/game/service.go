package game

import "regexp"

var pathRegex = regexp.MustCompile(`(\w{3})\s=\s\((\w{3}),\s(\w{3})\)`)

func FollowPath(instructions string, path map[string][]string) int {
	i := 0
	steps := 0

	currentSign := "AAA"

	for {
		steps++
		direction := instructions[i]

		if direction == 'R' {
			currentSign = path[currentSign][1]
		} else {
			currentSign = path[currentSign][0]
		}

		if currentSign == "ZZZ" {
			return steps
		}

		i++
		if i >= len(instructions) {
			i = 0
		}
	}
}

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
