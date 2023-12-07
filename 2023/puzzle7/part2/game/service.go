package game

import (
	"slices"
	"strconv"
	"strings"
)

type Service struct {
	input []string
}

func NewService(input []string) *Service {
	return &Service{
		input: input,
	}
}

func (s *Service) ToHands() []Hand {
	var result []Hand

	for _, line := range s.input {
		hand := toHand(line)
		hand.HandType = getHandType(hand)
		result = append(result, hand)
	}

	return result
}

func getHandType(hand Hand) HandType {
	result := HighCard

	occurrenceMap := map[string]int{}

	for _, card := range hand.Cards {
		occurrenceMap[card.CardType]++
	}

	lenMap := len(occurrenceMap)

	switch {
	case lenMap == 1:
		result = FiveOfAKind
	case lenMap == 2:
		for _, value := range occurrenceMap {
			if value == 4 {
				result = FourOfAKind
				break
			}

			if value == 3 || value == 2 {
				result = FullHouse
				break
			}
		}
	case lenMap == 3:
		for _, value := range occurrenceMap {
			if value == 3 {
				result = ThreeOfAKind
				break
			}

			if value == 2 {
				result = TwoPair
				break
			}
		}
	case lenMap == 4:
		result = OnePair
	case lenMap == 5:
		result = HighCard
	}

	result = upgradeJoker(result, hand)

	return result
}

func upgradeJoker(result HandType, hand Hand) HandType {
	jokerCount := 0
	for _, card := range hand.Cards {
		if card.CardType == "J" {
			jokerCount++
		}
	}

	if jokerCount == 0 || jokerCount == 5 {
		return result
	}

	newResult := result

	switch result {
	case HighCard:
		if jokerCount == 1 {
			newResult = OnePair
		}
	case OnePair:
		if jokerCount == 1 {
			newResult = ThreeOfAKind
		}

		if jokerCount == 2 {
			newResult = ThreeOfAKind
		}
	case TwoPair:
		if jokerCount == 1 {
			newResult = FullHouse
		}

		if jokerCount == 2 {
			newResult = FourOfAKind
		}
	case ThreeOfAKind:
		if jokerCount == 2 || jokerCount == 3 {
			newResult = FiveOfAKind
		} else {
			newResult = FourOfAKind
		}
	case FullHouse:
		if jokerCount == 2 {
			newResult = FiveOfAKind
		}

		if jokerCount == 3 {
			newResult = FiveOfAKind
		}
	case FourOfAKind:
		newResult = FiveOfAKind
	}

	println("Upgraded:", result.String(), "->", newResult.String(), "cards:", hand.CardString)
	return newResult
}

func toHand(line string) Hand {
	splitInput := strings.Split(line, " ")
	cardsString := splitInput[0]
	bidString := splitInput[1]

	bid, _ := strconv.Atoi(bidString)

	var cards []Card

	highCard := Card{}
	for _, cardSign := range cardsString {
		strength, _ := strconv.Atoi(string(cardSign))

		switch string(cardSign) {
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

		card := Card{Strength: strength, CardType: string(cardSign)}

		if strength > highCard.Strength {
			highCard = card
		}

		cards = append(cards, card)
	}

	result := Hand{
		Bid:         bid,
		Cards:       cards,
		HighestCard: highCard,
		CardString:  cardsString,
	}

	return result
}

func SortHands(hands []Hand) {
	slices.SortStableFunc(hands, sortHandleFunc)
}

func sortHandleFunc(i, j Hand) int {
	if i.HandType > j.HandType {
		return 1
	}

	if i.HandType < j.HandType {
		return -1
	}

	if i.HandType == j.HandType {
		for index, card1 := range i.Cards {
			card2 := j.Cards[index]
			if card1.Strength > card2.Strength {
				return 1
			}

			if card1.Strength < card2.Strength {
				return -1
			}
		}
	}

	return 0
}

func Play(hands []Hand) int {
	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.Bid
	}

	return result
}
