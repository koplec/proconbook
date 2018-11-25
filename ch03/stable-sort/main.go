package main

import "fmt"

type Card struct {
	Suit  string
	Value int
}

func main() {
	var n int
	var cards1, cards2 []Card

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	cards1 = make([]Card, n)
	for i := 0; i < n; i++ {
		card := Card{}
		_, err := fmt.Scanf("%s%d", &card.Suit, &card.Value)
		if err != nil {
			panic(err)
		}
		cards1[i] = card
	}

	for i := 0; i < n; i++ {
		cards1[i] = cards2[i]
	}

	bubble(cards1)
	selection(cards2)

	print(cards1)
	fmt.Printf("Stable\n")

	print(cards2)
	if isStable(cards1, cards2) {
		fmt.Printf("Stable\n")
	} else {
		fmt.Printf("Not stable\n")
	}
}

func bubble(cards []Card) { //bubble sort is stable
	n := len(cards)
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if cards[j].Value < cards[j-1].Value {
				cards[j], cards[j-1] = cards[j-1], cards[j]
			}
		}
	}
}

func selection(cards []Card) {
	n := len(cards)
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if cards[j].Value < cards[minj].Value {
				minj = j
			}
		}
		cards[minj], cards[i] = cards[i], cards[minj]
	}
}

func print(cards []Card) {
	for i := 0; i < len(cards); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%s%d", cards[i].Suit, cards[i].Value)
	}
}

func isStable(cards1 []Card, cards2 []Card) bool {
	n := len(cards1)
	for i := 0; i < n; i++ {
		if cards1[i].Suit != cards2[i].Suit {
			return false
		}
	}
	return true
}
