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
		"From const Jack of Clubs": {
			rank: card.Jack, val: 11, suit: card.Clubs,
		},
		"16 of Diamonds": {
			rank: 16, val: 16, suit: card.Dimaonds,
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
				t.Errorf("Got\n\t%+v\nwanted\n\t%+v\n", got, want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := map[string]struct {
		c    card.Card
		want string
	}{
		"Ace of Spades": {
			c:    card.New(14, 11, card.Spades),
			want: "\u2660Ace",
		},
		"Queen of Hearts": {
			c:    card.New(card.Queen, 10, card.Hearts),
			want: "\u2665Queen",
		},
		"10 of Clubs": {
			c:    card.New(10, 10, card.Clubs),
			want: "\u266310",
		},
		"6 of Diamonds": {
			c:    card.New(6, 6, card.Dimaonds),
			want: "\u26666",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.c.String()
			if want := tc.want; got != want {
				t.Errorf("Got %q, wanted %q", got, want)
			}
		})
	}
}
