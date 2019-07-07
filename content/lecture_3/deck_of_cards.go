package main

import (
	"fmt"
	"math/rand"
)

func main() {
	deck := Deck{
		cards: []Card{
			{"Туз", "Пики"},
			{"Дама", "Треф"},
			{"Король", "Буби"},
			{"Шестерка", "Черви"},
		},
	}

	fmt.Println(deck)

	deck.shuffle()

	fmt.Println(deck)
}

type Card struct {
	value  string
	colour string
}

type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(len(d.cards) - 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
