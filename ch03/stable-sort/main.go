package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	Suit  string
	Value int
}

func main() {
	var n int
	var cards1 []Card
	var cards2 []Card

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("n = %d\n", n)

	cards1 = make([]Card, n)
	cards2 = make([]Card, n)

	for i := 0; i < n; i++ {
		card := Card{}
		var cardStr string
		_, err := fmt.Scan(&cardStr)
		//fmt.Scanf("%s", &cardStr)としたらunexpected new line exceptionが発生　なんで？
		if err != nil {
			fmt.Printf("error occured %v\n", err)
			panic(err)
		}
		//fmt.Printf("%s\n", cardStr)
		for j, s := range cardStr {
			if j == 0 {
				card.Suit = string(s)
			} else {
				card.Value, _ = strconv.Atoi(string(s))
			}
		}
		cards1[i] = card

		fmt.Printf("cards1[%d] = %v\n", i, cards1[i])
	}

	fmt.Printf("copy 1 to 2\n")
	for i := 0; i < n; i++ {
		fmt.Printf("i=%d\n", i)
		cards2[i] = cards1[i]
	}

	fmt.Printf("cards1 -> %v\n", cards1)
	fmt.Printf("cards2 -> %v\n", cards2)

	fmt.Printf("bubble sort start\n")
	bubble(cards1)
	fmt.Printf("selection sort start\n")
	selection(cards2)
	fmt.Printf("after sort\n")
	fmt.Printf("cards1 -> %v\n", cards1)
	fmt.Printf("cards2 -> %v\n", cards2)
	//bubble sort is stable
	if isStable(cards1, cards2) {
		fmt.Printf("Stable\n")
	} else {
		fmt.Printf("Not stable\n")
	}
}

func bubble(cards []Card) { //bubble sort is stable
	n := len(cards)
	for i := 1; i < n; i++ {
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
