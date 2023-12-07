package game

type HandType int

// HandTypes with their values.
const (
	HighCard     HandType = iota + 1
	OnePair      HandType = iota + 1
	TwoPair      HandType = iota + 1
	ThreeOfAKind HandType = iota + 1
	FullHouse    HandType = iota + 1
	FourOfAKind  HandType = iota + 1
	FiveOfAKind  HandType = iota + 1
)

type Hand struct {
	Cards       []Card
	CardString  string
	HighestCard Card
	Bid         int
	HandType    HandType
	Strength    int
}

type Card struct {
	Strength int
	CardType string
}
