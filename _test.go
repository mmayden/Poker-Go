// draft, doesn't work
package main

import "testing"

func TestNewTable(t *testing.T) {
	playas := newPlayers()
	fmt.Println(playas)
	t.Errorf("What happened")
}
