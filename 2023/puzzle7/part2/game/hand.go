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

func (h HandType) String() string {
	switch h {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	default:
		return "Unknown"
	}
}

type Hand struct {
	Cards       []Card
	CardString  string
	HighestCard Card
	Bid         int
	HandType    HandType
}

type Card struct {
	Strength int
	CardType string
}
