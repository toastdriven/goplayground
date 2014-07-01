package deck

import (
	"math/rand"
)

var Values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
const Club = "♧"
const Spade = "♤"
const Heart = "♡"
const Diamond = "♢"
var Suits = []string{
	Club,
	Spade,
	Heart,
	Diamond,
}

type Card struct {
	value string
	suit  string
}

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := Deck{}

	for i := range Values {
		for j := range Suits {
			card := Card{Values[i], Suits[j]}
			deck.cards = append(deck.cards, card)
		}
	}

	return &deck
}

func (deck *Deck) Draw() Card {
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card
}

func (deck *Deck) ReturnBottom(card Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) RandomCard() Card {
	// This isn't particularly efficient (for now), since there's lots
	// of new slice/array creations going on.
	offset := rand.Intn(len(deck.cards))
	card := deck.cards[offset]
	deck.cards = append(deck.cards[:offset], deck.cards[offset+1:]...)
	return card
}

func (deck *Deck) Shuffle() {
	shuffled := []Card{}

	for len(deck.cards) > 0 {
		random_card := deck.RandomCard()
		shuffled = append(shuffled, random_card)
	}

	deck.cards = shuffled
}
