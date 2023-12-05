package main

import "fmt"

//main game loop
func main() {
	if err := playpokergo(); err != nil {
		fmt.Println(err)
	}
}

//main game table logic loop
func playpokergo() error {

	// create table
	pokergo_table := table{newPlayers(), shuffle(), 0, 0, 0, nil}
	// set player names
	for i := 0; i < len(pokergo_table.players); i++ {
		pokergo_table.setPlayer(i, getName(i))
	}

	//game loop
	for {
		//deal table
		pokergo_table.deal()
		//table plays hand
		pokergo_table.playhand()
		//update table at end of hand
		//pokgero_table.update()
	}
}