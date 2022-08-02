package main

import (
	"fmt"
	"strconv"
)

type player struct {
	name    string
	hand    []string
	handVal int
	blind   int
}

func (tablePointer *table) addPlayer() {
	tablePointer.players = append(tablePointer.players, player{})
}

func (tablePointer *table) rmPlayer() {
	copy(tablePointer.players[2:], tablePointer.players[3:])
	tablePointer.players = tablePointer.players[:len(tablePointer.players)-1]
}

func (tablePointer *table) setPlayer(n int, s string) {
	tablePointer.players[n].name = s
}

func getName(i int) string {
	var r, r2 string
	for {
		fmt.Println("What is player " + strconv.Itoa(i+1) + "'s name")
		fmt.Scanln(&r)
		fmt.Println("Is " + r + " correct?")
		fmt.Scanln(&r2)
		if r2 == "y" || r2 == "yes" {
			return r
		}
	}
}

func newPlayers() []player {

	// return int # of players
	var numPlayers = func() int {
		var r, r2 string
		for {
			fmt.Println("How many players?")
			fmt.Scanln(&r)
			fmt.Println("Is " + r + " correct?")
			fmt.Scanln(&r2)
			if r2 == "y" || r2 == "yes" {
				np, err := strconv.Atoi(r)
				if err == nil {
					return np
				}
			}
		}
	}

	return make([]player, numPlayers())
}
