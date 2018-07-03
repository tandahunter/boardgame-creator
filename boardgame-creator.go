package main

import (
	"fmt"
	"math/rand"
)

//Broad game definitions
const (
	singlePlayer = iota
	multiPlayer
	singleDeck
	multiDeck
)

//Card transformations
const (
	rotate90 = iota
	rotate180
	rotate270
	flipHorizontal
)

//GameContext is the standard game definition class
type GameContext struct {
	definition int
}

//Point represents a 2d location
type Point struct {
	x int
	y int
}

//Deck contains multiple cards
type Deck struct {
	cards    *[]Card
	rotation int
	position *Point
}

//MoveTopToBottom moves the top card of the deck to the bottom
func (deck *Deck) MoveTopToBottom() {
	cards := *deck.cards
	lastIndex := len(cards) - 1

	var topCard Card

	for i, b := range cards {
		if i == 0 {
			topCard = b
		} else if i == lastIndex {
			cards[i] = topCard
		} else {
			cards[i-1] = b
		}
	}
}

//Shuffle shuffles the cards in the deck
func (deck *Deck) Shuffle() {
	cards := *deck.cards
	max := len(cards) - 1

	for i := range cards {
		cards[i] = cards[rand.Intn(max)]
	}
}

//Card contains 2 sides
type Card struct {
	id       int
	front    []byte
	back     []byte
	actions  []int
	rotation int
	isFront  bool
	position *Point
}

//Transform manipulates the rotation of facing of the card
func (card *Card) Transform(transformation int) {
	switch transformation {
	case rotate90:
		card.rotation = (card.rotation + 90) % 360
	case rotate180:
		card.rotation = (card.rotation + 180) % 360
	case rotate270:
		card.rotation = (card.rotation + 270) % 360
	case flipHorizontal:
		card.isFront = !card.isFront
		card.rotation = (card.rotation + 180) % 360
	}
}

func main() {
	fmt.Println("Hello, World")
}
