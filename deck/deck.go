package deck

import (
	"card"
	"math/rand"
)

type Deck struct {
	backingDeck    []card.Card
	drawPile       chan card.Card
	discardPile    []card.Card
	drawIterations int
	rankVal        map[int]int
}

func New(options ...Option) (d Deck) {
	d.drawIterations = 1
	d.drawPile = make(chan card.Card)

	for _, option := range options {
		option(&d)
	}

	rv := &(d.rankVal)
	if *rv == nil {
		rv = &map[int]int{
			1:  11,
			11: 10,
			12: 10,
			13: 10,
		}
	}

	d.backingDeck = backingDeck(*rv)
	db := d.backingDeck
	dd := d.drawIterations
	d.Shuffle()

	go func() {
		defer close(d.drawPile)
		for dd != 0 {
			for _, card := range db {
				d.drawPile <- card
			}

			dd--
		}
	}()

	return
}

func backingDeck(rankVal map[int]int) []card.Card {
	cards := make([]card.Card, 52)
	var i int
	for s := range [4]struct{}{} {
		for r := 1; r < 14; r++ {
			v, ok := rankVal[r]
			if !ok {
				v = r
			}
			cards[i] = card.New(r, v, card.Suit(s))
			i++
		}
	}
	return cards
}

func (d Deck) DrawPile() chan card.Card {
	return d.drawPile
}

func (d Deck) DiscardPile() []card.Card {
	return d.discardPile
}

func (d Deck) Shuffle() {
	b := d.backingDeck
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
}

type Option func(*Deck)

func DrawIterations(n int) Option {
	return func(d *Deck) {
		d.drawIterations = n
	}
}

func RankValueMap(m map[int]int) Option {
	return func(d *Deck) {
		d.rankVal = m
	}
}
