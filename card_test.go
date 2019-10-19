package card_test

import (
	"card"
	"testing"
)

func TestNew(t *testing.T) {
	type wanted struct {
		rank, val int
		suit      card.Suit
	}
	tests := map[string]wanted{
		"Ace of Spades": {
			rank: 14, val: 11, suit: card.Spades,
		},
		"2 of Hearts": {
			rank: 2, val: 2, suit: card.Hearts,
		},
	}
	equal := func(c card.Card, t wanted) bool {
		return c.Rank() == t.rank &&
			c.Value() == t.val &&
			c.Suit() == t.suit
	}
	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			got := card.New(want.rank, want.val, want.suit)
			if !equal(got, want) {
				t.Errorf("%s: got\n\t%+v\nwanted\n\t%+v\n", name, got, want)
			}
		})
	}
}
