package game_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle7/part2/game"
	"github.com/stretchr/testify/require"
)

func TestInput2(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	gameService := game.NewService(input)
	hands := gameService.ToHands()

	game.SortHands(hands)

	result := game.Play(hands)

	require.Equal(t, 5905, result)

	fmt.Println(hands)
}

func TestSingleHand(t *testing.T) {
	t.Run("OnePair", func(t *testing.T) {
		input := []string{
			"32T3K 765",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.OnePair)
	})

	t.Run("ThreeOfAKind", func(t *testing.T) {
		input := []string{
			"T55J5 684",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.FourOfAKind)
	})

	t.Run("TwoPair", func(t *testing.T) {
		input := []string{
			"KK677 28",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.TwoPair)
	})

	t.Run("TwoPair #2", func(t *testing.T) {
		input := []string{
			"KTJJT 220",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.FourOfAKind)
	})

	t.Run("ThreeOfAKind", func(t *testing.T) {
		input := []string{
			"QQQJA 483",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.FourOfAKind)
	})

	t.Run("4OfAKind", func(t *testing.T) {
		input := []string{
			"TTT4T 627",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.FourOfAKind)
	})

}

func TestSortHands(t *testing.T) {
	hands := []game.Hand{
		{
			Cards: []game.Card{
				{
					CardType: "K",
					Strength: getCardStrength("K"),
				},
				{
					CardType: "2",
				},
				{
					CardType: "3",
				},
				{
					CardType: "4",
				},
				{
					CardType: "5",
				},
			},
		},
		{
			Cards: []game.Card{
				{
					CardType: "J",
					Strength: getCardStrength("J"),
				},
				{
					CardType: "2",
				},
				{
					CardType: "3",
				},
				{
					CardType: "4",
				},
				{
					CardType: "5",
				},
			},
		},
		{
			Cards: []game.Card{
				{
					CardType: "T",
					Strength: getCardStrength("T"),
				},
				{
					CardType: "2",
				},
				{
					CardType: "3",
				},
				{
					CardType: "4",
				},
				{
					CardType: "5",
				},
			},
		},
	}

	game.SortHands(hands)

	require.Equal(t, hands[0].Cards[0].CardType, "J")
}

func getCardStrength(card string) int {
	strength, _ := strconv.Atoi(string(card))

	switch string(card) {
	case "T":
		strength = 10
	case "J":
		strength = 1
	case "Q":
		strength = 12
	case "K":
		strength = 13
	case "A":
		strength = 14
	}

	return strength
}
