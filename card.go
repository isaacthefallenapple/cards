package card

import "fmt"

type Suit int

const (
	Clubs Suit = iota
	Dimaonds
	Hearts
	Spades
)

const (
	Ace  = 1
	Jack = iota + 11
	Queen
	King
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

func (c Card) String() string {
	suits := [4]rune{'\u2663', '\u2666', '\u2665', '\u2660'}
	var name string
	switch r := c.rank; r {
	case 1:
		name = "Ace"
	case 13:
		name = "King"
	case 12:
		name = "Queen"
	case 11:
		name = "Jack"
	default:
		name = fmt.Sprint(r)
	}
	return fmt.Sprintf("%c%s", suits[c.suit], name)
}
