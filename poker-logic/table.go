package main

import (
	"fmt"
)

type table struct {
	players []player
	dealer  deck
	round   string //used for dealer/smblind/bblind
}

func (tablePointer *table) deal() {

	// set hands empty
	for i := 0; i < len(tablePointer.players); i++ {
		tablePointer.players[i].hand = make([]string, 2)
	}

	// deal 1 card to each player twice
	var counter int
	for i := 0; i < 2; i++ {
		for j := 0; j < len(tablePointer.players); j++ {
			tablePointer.players[j].hand[i] += tablePointer.dealer[counter]
			//remove dealt cards from deck
			tablePointer.dealer = tablePointer.dealer[1:]
			counter++
		}
	}
}

func (tablePointer *table) playhand() {
	for i := 0; i < len(tablePointer.players); i++ {
		fmt.Println("Bet is to you " + tablePointer.players[i].name)
		bet(tablePointer.players[i])

	}
}

func play() error {

	//create table
	t := table{newPlayers(), shuffle(), 0}
	//set player names
	for i := 0; i < len(t.players); i++ {
		fmt.Println("Before")
		t.setPlayer(i, getName(i))
		fmt.Println("After")
	}

	//game loop
	for {
		t.deal()
		t.playhand()
		/*
			for i := 0; i < len(t.players); i++ {
				fmt.Println(t.players[i].hand)
			}
			fmt.Println(len(t.dealer))
		*/
		return nil
	}
}

//fmt.Println("There are " + strconv.Itoa(len(t.players)) + " gamblers with us here.")
