package main

import (
	"fmt"
)

type table struct {
	players []player
	dealer  deck
	//bigblind	string //used for dealer/smblind/bblind
	pot      int
	topbet   int
	betround int
	board    []string
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

	for {
		// flop
		if tablePointer.betround == 1 {
			for i := 0; i < 3; i++ {
				tablePointer.board = append(tablePointer.board, tablePointer.dealer[0])
				tablePointer.dealer = tablePointer.dealer[1:]
			}
			fmt.Println("The flop is    ", tablePointer.board)
		}

		// turn/river
		if tablePointer.betround == 2 || tablePointer.betround == 3 {
			tablePointer.board = append(tablePointer.board, tablePointer.dealer[0])
			tablePointer.dealer = tablePointer.dealer[1:]
			fmt.Println("Next card:    ", tablePointer.board)
		}

		// showHands()
		if tablePointer.betround == 3 {
			tablePointer.betround = 0
			for k := 0; k < len(tablePointer.players); k++ {
				fmt.Println(tablePointer.players[k].name, "s hand:    ", tablePointer.players[k].hand)
			}

		}

		// next player
		for j := 0; j < len(tablePointer.players); j++ {
			fmt.Println("Bet is to you ", tablePointer.players[j].name)
			tablePointer.players[j].bet()
			//if type number, bet, else they check/fold
			//if (tablePointer.players[j].currentBet )
			fmt.Println(tablePointer.players[j].name + " bet " + tablePointer.players[j].currentBet + "\n")
		}
		tablePointer.betround++
	}
}

func play() error {

	// create table
	t := table{newPlayers(), shuffle(), 0, 0, 0, nil}
	// set player names
	for i := 0; i < len(t.players); i++ {
		t.setPlayer(i, getName(i))
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