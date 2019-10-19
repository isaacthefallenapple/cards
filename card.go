package card

type suit uint8

const (
	Clubs suit = iota
	Dimaonds
	Hearts
	Spades
)

type Card struct {
	rank, value int
	suit        suit
}

func New(rank, value int, suit suit) Card {
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

func (c Card) Suit() suit {
	return c.suit
}
