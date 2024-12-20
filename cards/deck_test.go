package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand length of 5, but got %v", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected remaining deck length of 47, but got %v", len(remainingDeck))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	if d[0] == "Ace of Spades" && d[1] == "Two of Spades" && d[2] == "Three of Spades" && d[3] == "Four of Spades" {
		t.Errorf("Expected deck to be shuffled, but it's not")
	}
}
