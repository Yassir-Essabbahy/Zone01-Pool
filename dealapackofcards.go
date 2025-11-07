package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	nbrPlayers := 4
	Cards := 3
	for i := 0; i < nbrPlayers; i++ {
		s := i * Cards
		e := s + Cards
		d := deck[s:e]
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, d[0], d[1], d[2])
	}
}
