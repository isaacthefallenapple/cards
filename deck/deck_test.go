package deck_test

import (
	"card"
	"card/deck"
	"testing"
)

func TestNew(t *testing.T) {
	d := deck.New()
	var (
		ok bool
		c  card.Card
	)
	for range [53]struct{}{} {
		c, ok = <-d.DrawPile()
		t.Log(c, ok)
	}
	if !ok {
		t.Error("There shouldn't be any more cards")
		t.Logf("Next: %s", <-d.DrawPile())
	}
}
