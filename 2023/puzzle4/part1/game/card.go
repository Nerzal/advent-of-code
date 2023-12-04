package game

import "slices"

type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
}

func Solve(cards []Card) int {
	result := 0
	for _, card := range cards {
		cardSum := 0
		for _, number := range card.Numbers {
			if !slices.Contains(card.WinningNumbers, number) {
				continue
			}

			if cardSum == 0 {
				cardSum = 1
				continue
			}

			cardSum *= 2
		}
		result += cardSum
	}

	return result
}
