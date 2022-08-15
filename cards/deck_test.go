package main

import (
	"os"
	"testing"
)

func TestNewDect(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected d[0] to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected d[0] to be 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	deckFromFile := newDeckFromFile(filename)

	if len(deckFromFile) != 16 {
		t.Errorf("Expected to have 16 items in deck loaded from file, but got %v", len(deckFromFile))
	}

	os.Remove(filename)
}
