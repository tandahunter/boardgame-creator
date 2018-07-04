package main

import (
	"fmt"
	"math/rand"
	"os"
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

//Area definitions
const (
	topHalf = iota
	bottomHalf
	front
	back
)

//Currency represents a currency item
type Currency struct {
	id   int
	name string
}

//GameContext is the standard game definition class
type GameContext struct {
	name       string
	definition int
	decks      *[]Deck
}

//Point represents a 2d location
type Point struct {
	x int
	y int
}

//CardArea represents an area on a card
type CardArea struct {
	area    int
	value   *[]Currency
	actions *[]CardAction
	points  int
}

//CardAction represents an action performed by a CardArea
type CardAction struct {
	action int
	cost   *[]Currency
}

//Card contains 2 sides
type Card struct {
	id       int
	name     string
	rotation int
	front    []byte
	back     []byte
	isFront  bool
	position *Point
	areas    *[]CardArea
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
		card, newIndex := cards[i], rand.Intn(max)

		cards[i] = cards[newIndex]
		cards[newIndex] = card
	}
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

func loadContext(filename string) {
	xmlFile, err := os.Open("data/palm-island.cfg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()

}
