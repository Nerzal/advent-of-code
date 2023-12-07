package game_test

import (
	"fmt"
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game"
	"github.com/stretchr/testify/assert"
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

	require.Equal(t, 6440, result)

	fmt.Println(hands)
}

func TestInput3(t *testing.T) {
	input := []string{
		"4K5T5 643", // One Pair #3
		"97665 993", // One Pair #5
		"JJ52T 822", // One Pair
		"93898 538", // One Pair
		"74944 966", // ThreeOfAKind
		"3K33K 1",   // FullHouse
		"JT8TQ 15",  // 2 Pairs #7
		"7T777 327", // 4OfAKind #14
		"483A2 836", // HighCard #1
		"2828J 960", // 2 Pairs
		"J7K79 434", // One Pair #6
		"54J5T 979", // 1 Pair #4
		"6ATKJ 552", // HighCard #2
		"TTT4T 627", // 4OfAKind #15
		"6566J 598", // ThreeOfAKind
	}

	gameService := game.NewService(input)
	hands := gameService.ToHands()

	game.SortHands(hands)

	result := game.Play(hands)

	if !assert.Equal(t, result, 68221) {
		println("fail")
	}
	require.Equal(t, "483A2", hands[0].CardString)
	require.Equal(t, "6ATKJ", hands[1].CardString)
	require.Equal(t, "4K5T5", hands[2].CardString)
	t.Log(result)
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

		require.Equal(t, hands[0].HandType, game.ThreeOfAKind)
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

		require.Equal(t, hands[0].HandType, game.TwoPair)
	})

	t.Run("ThreeOfAKind", func(t *testing.T) {
		input := []string{
			"QQQJA 483",
		}

		gameService := game.NewService(input)
		hands := gameService.ToHands()

		require.Equal(t, hands[0].HandType, game.ThreeOfAKind)
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
					Strength: game.GetCardStrength("K"),
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
					Strength: game.GetCardStrength("J"),
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
					Strength: game.GetCardStrength("T"),
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

	require.Equal(t, hands[0].Cards[0].CardType, "T")
}
