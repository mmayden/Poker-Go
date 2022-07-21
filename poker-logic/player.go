package main

import (
	"fmt"
	"strconv"
)

type player struct {
	name    string
	hand    []string
	handVal int
}

func (tablePlayersPointer *table) addPlayer() {
	tablePlayersPointer.players = append(tablePlayersPointer.players, player{})
}

func (tablePlayersPointer *table) rmPlayer() {
	copy(tablePlayersPointer.players[2:], tablePlayersPointer.players[3:])
	tablePlayersPointer.players = tablePlayersPointer.players[:len(tablePlayersPointer.players)-1]
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

	//create player slice and init each player
	return make([]player, (numPlayers()))
}
