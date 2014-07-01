package deck

import (
	"testing"
	"math/rand"
)

func TestDeck(t *testing.T) {
	// For consistency in testing...
	rand.Seed(4)

	deck := NewDeck()

	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards, was:", len(deck.cards))
	}

	top_card := deck.Draw()

	if top_card.value != "2" {
		t.Errorf("Expected top card of 2, was:", top_card.value)
	}

	if top_card.suit != Club {
		t.Errorf("Expected top card of clubs, was:", top_card.suit)
	}

	deck.ReturnBottom(top_card)

	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards, was:", len(deck.cards))
	}

	next_card := deck.Draw()

	if next_card.value != "2" {
		t.Errorf("Expected next card of 2, was:", next_card.value)
	}

	if next_card.suit != Spade {
		t.Errorf("Expected next card of clubs, was:", next_card.suit)
	}

	deck.ReturnBottom(next_card)

	// Go through more of the deck.
	for i := 0; i < 13; i++ {
		next_card = deck.Draw()
		deck.ReturnBottom(next_card)
	}

	next_card = deck.Draw()

	if next_card.value != "5" {
		t.Errorf("Expected next card of 5, was:", next_card.value)
	}

	if next_card.suit != Diamond {
		t.Errorf("Expected next card of diamonds, was:", next_card.suit)
	}

	deck.ReturnBottom(next_card)

	// Pull a random card.
	random_card := deck.RandomCard()

	if random_card.value != "9" {
		t.Errorf("Expected random card of 9, was:", random_card.value)
	}

	if random_card.suit != Spade {
		t.Errorf("Expected random card of spades, was:", random_card.suit)
	}

	deck.ReturnBottom(random_card)

	// Now let's try shuffling.
	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards pre-shuffle, was:", len(deck.cards))
	}

	deck.Shuffle()

	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards post-shuffle, was:", len(deck.cards))
	}

	top_shuffled_card := deck.Draw()

	if top_shuffled_card.value != "3" {
		t.Errorf("Expected random card of 3, was:", top_shuffled_card.value)
	}

	if top_shuffled_card.suit != Spade {
		t.Errorf("Expected random card of spades, was:", top_shuffled_card.suit)
	}

	deck.ReturnBottom(top_shuffled_card)
	top_shuffled_card = deck.Draw()

	if top_shuffled_card.value != "A" {
		t.Errorf("Expected random card of A, was:", top_shuffled_card.value)
	}

	if top_shuffled_card.suit != Diamond {
		t.Errorf("Expected random card of diamonds, was:", top_shuffled_card.suit)
	}
}
