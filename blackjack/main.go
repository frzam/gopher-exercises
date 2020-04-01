package main

import (
	"fmt"
	"strings"

	"github.com/farzamalam/gopher-exercises/deck"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player : ", player)
		fmt.Println("Dealer : ", dealer.dealerString())
		fmt.Println("What will you do? (h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	// If dealer's score is <= 16, we hit
	// If a dealer has soft 17 then we hit.

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.minScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("-------- Final Hand ---------")
	fmt.Println("Dealer : ", dealer)
	fmt.Println("Dealer score : ", dScore)
	fmt.Println("Player : ", player)
	fmt.Println("Player score : ", pScore)

	switch {
	case pScore > 21:
		fmt.Println("You are busted!")
	case dScore > 21:
		fmt.Println("Dealer is busted!")
	case pScore > dScore:
		fmt.Println("You win !!!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
}

func (h Hand) Score() int {
	minScore := h.minScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}
func (h Hand) minScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	card, cards := cards[0], cards[1:]
	return card, cards
}

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) dealerString() string {
	return h[0].String() + ", ** Hidden **"
}
