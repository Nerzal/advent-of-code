package game

import "slices"

type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
}

type GameData struct {
	Cards []Card
}

func NewGameData(cards []Card) GameData {
	return GameData{
		Cards: cards,
	}
}

func (g GameData) Solve(cards []Card) int {
	result := 0
	for _, card := range cards {
		copies := g.getCopies(card)
		result += len(copies)

		result += g.Solve(copies)
	}

	return result
}

func (g GameData) getCopies(card Card) []Card {
	var copies []Card

	copyIndex := 1
	for _, number := range card.Numbers {
		if !slices.Contains(card.WinningNumbers, number) {
			continue
		}

		copies = append(copies, g.Cards[card.ID+copyIndex-1])
		copyIndex++
	}

	return copies
}
