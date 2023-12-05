package main

import (
	"fmt"
	"strconv"
)

type player struct {
	name       string
	hand       []string
	handVal    int
	blind      int
	chips      int
	currentBet string
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
	tablePointer.players[n].chips = 1000
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

func (playerPointer *player) bet() {
	//r = user return used below, b = bet decision that is returned
	var r int
	var b string
	//var getBet = func() string {
	for {
		fmt.Println("1. \"Check\"\n2. \"Bet\"\n3. \"Call\"\n4. \"Fold\"")
		fmt.Scanln(&r)
		if r == 1 {
			playerPointer.currentBet = "check"
			break
		} else if r == 2 {
			for {
				fmt.Println("How much?")
				fmt.Scanln(&b)
				//convert string to int
				i, err := strconv.Atoi(b)
				if err != nil {
					fmt.Println("Not a numeric value..")
				} else if i > playerPointer.chips {
					fmt.Println("You ain't got the coin..")
				} else if i < 1 {
					fmt.Println("Quit playin..")
				} else {
					playerPointer.currentBet = b
					fmt.Printf(playerPointer.currentBet + " whats wrong[%T]\nYet we got " + b)
					break
				}
			}
		} else if r == 3 {
			playerPointer.currentBet = "call"
			break
		} else if r == 4 {
			playerPointer.currentBet = "fold"
			break
		} else {
			fmt.Println("Invalid response.")
		}
	}
	return
}

func handVal() {
	
}
