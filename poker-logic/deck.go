// deck

package main

import (
	"math/rand"
	"time"
)

type deck []string

// create deck and shuffle, return shuffled deck
func shuffle() deck {
	// create deck
	d := deck{}

	// slices of card suits and values
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardVals := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	// add cardval of each suit to deck
	for _, suit := range suits {
		for _, cardVal := range cardVals {
			d = append(d, cardVal+" of "+suit)
		}
	}

	// shuffle deck
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
