// main

package main

import (
	"fmt"
)

func main() {
	if err := play(); err != nil {
		fmt.Println(err)
	}
}
