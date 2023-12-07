package main

import (
	"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game"
	"github.com/Nerzal/advent-of-code/2023/template/file"
)

func main() {
	input, err := file.GetInputData()
	if err != nil {
		panic(err)
	}

	gameService := game.NewService(input)
	hands := gameService.ToHands()

	game.SortHands(hands)

	result := game.Play(hands)
	println(result)
	// 248965970
	// 249021152
	// 248887050
	// 248855457

	//game.Hand {Cards: []github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card len: 5, cap: 8, [(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba40),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba58),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba70),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba88),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018baa0)], HighestCard: github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card {Strength: 12, CardType: "Q"}, Bid: 269, HandType: TwoPair (3), Strength: 34}
	//game.Hand {Cards: []github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card len: 5, cap: 8, [(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba40),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba58),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba70),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba88),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018baa0)], HighestCard: github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card {Strength: 12, CardType: "Q"}, Bid: 269, HandType: TwoPair (3), Strength: 34}
	//game.Hand {Cards: []github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card len: 5, cap: 8, [(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba40),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba58),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba70),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018ba88),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc00018baa0)], HighestCard: github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card {Strength: 12, CardType: "Q"}, Bid: 269, HandType: TwoPair (3), Strength: 34}

	//game.Hand {Cards: []github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card len: 5, cap: 8, [(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001edb00),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001edb18),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001edb30),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001edb48),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001edb60)], HighestCard: github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card {Strength: 13, CardType: "K"}, Bid: 497, HandType: FullHouse (5), Strength: 63}
	//game.Hand {Cards: []github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card len: 5, cap: 8, [(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001c78c0),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001c78d8),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001c78f0),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001c7908),(*"github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card")(0xc0001c7920)], HighestCard: github.com/Nerzal/advent-of-code/2023/puzzle7/part1/game.Card {Strength: 12, CardType: "Q"}, Bid: 939, HandType: FullHouse (5), Strength: 40}

}
