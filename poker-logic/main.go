// main

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// create table
	t := table{newPlayers(), shuffle()}
	fmt.Println("There are " + strconv.Itoa(len(t.players)) + " gamblers with us here.")
	play(t)
}

type table struct {
	players []player
	dealer  deck
}

type player struct {
	name    string
	hand    []string
	handVal int
}

func (tablePointer *table) tableinit() {
	//player inits
}

func (playerPointer *player) playerinit() {
	// playername
	// (if game is holdem) set handsize to 2
	playerPointer.hand = make([]string, 2)
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
	players := make([]player, numPlayers())
	return players
}

func (tablePointer *table) deal() {
	// deal 1 card to each player twice
	var counter int
	for i := 0; i < 2; i++ {
		for j := 0; j < len(tablePointer.players); j++ {
			tablePointer.players[j].hand[i] += tablePointer.dealer[counter]
			counter++
		}
	}
	// remove dealt cards from deck
}

func play(t table) {
	for {
		t.deal()
		fmt.Println(t.dealer)
		return
	}
}
