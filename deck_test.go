package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

//TestMoveTopToBottom tests the Deck.MoveTopToBottom routine
func TestMoveTopToBottom(t *testing.T) {
	deck := Deck{}
	top, middle, bottom := Card{id: 1}, Card{id: 2}, Card{id: 3}
	deck.cards = &[]Card{top, middle, bottom}
	cards := *deck.cards

	if cards[0].id != top.id {
		t.Error("Initial Top card is wrong")
	}

	deck.MoveTopToBottom()

	log.Printf(strconv.Itoa(cards[0].id))

	if cards[0].id == top.id {
		t.Error("The Top card is still on top")
	}

	if cards[2].id != top.id {
		t.Error("The Top card has not been moved to the bottom")
	}
}

//TestShuffle tests the Deck.Shuffle routine
func TestShuffle(t *testing.T) {
	deckcount := 100
	offset := 1

	deck := Deck{}
	cards := make([]Card, deckcount)
	deck.cards = &cards

	for i := 0; i < deckcount; i++ {
		cards[i] = Card{id: i + offset}
	}

	for i := 0; i < deckcount; i++ {
		fmt.Println(strconv.Itoa(cards[i].id))
	}

	deck.Shuffle()

	avg := (float32(offset+deckcount) / 2)
	desiredtotal := int(avg * (float32(deckcount)))

	decktotal := 0
	for i := 0; i < deckcount; i++ {
		decktotal += cards[i].id
	}

	if desiredtotal != decktotal {
		t.Error("The deck is corrupt, expected " + strconv.Itoa(desiredtotal) + " received " + strconv.Itoa(decktotal))
	}
}
