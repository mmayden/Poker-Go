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

//pokergo game table step 1. deal hand
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

//pokergo game table step 2. playhand
func (tablePointer *table) playhand() {
	//playhand
	for {
		//hand is announced
		if tablePointer.betround == 1 {
			fmt.Println("The flop is    ", tablePointer.board)
		} else if tablePointer.betround == 2 {
			fmt.Println("Turn Card:    ", tablePointer.board)
		}else if tablePointer.betround == 3 {
			fmt.Println("River Card:    ", tablePointer.board)
		}
			
		//player bets	
		for j := 0; j < len(tablePointer.players); j++ {
			//Player plays
			fmt.Println("Bet is to you ", tablePointer.players[j].name)
			//Player bets
			tablePointer.players[j].bet()
			//if type number, bet, else they check/fold
			//if (tablePointer.players[j].currentBet )
			//Announce player's bet
			fmt.Println(tablePointer.players[j].name + " bet " + tablePointer.players[j].currentBet + "\n")
		}

		tablePointer.betround++
		//update board
		tablePointer.board = append(tablePointer.board, tablePointer.dealer[0])
		//update dealer
		tablePointer.dealer = tablePointer.dealer[1:]

		//hand is over, show hands, show winner, reset betround to 0
		if tablePointer.betround == 4 {
			for k := 0; k < len(tablePointer.players); k++ {
				//show hands
				fmt.Println(tablePointer.players[k].name, "s hand:    ", tablePointer.players[k].hand)
			}
			//show winner
			fmt.Println("Coming Soon..")
			return
		}
	}
}