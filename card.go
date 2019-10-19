package card

type Suit uint8

const (
	Clubs Suit = iota
	Dimaonds
	Hearts
	Spades
)

type Card struct {
	rank, value int
	suit        Suit
}

func New(rank, value int, suit Suit) Card {
	return Card{
		rank:  rank,
		value: value,
		suit:  suit,
	}
}

func (c Card) Rank() int {
	return c.rank
}

func (c Card) Value() int {
	return c.value
}

func (c Card) Suit() Suit {
	return c.suit
}
