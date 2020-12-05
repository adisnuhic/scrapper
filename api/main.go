package main

import (
	"fmt"
	"time"
)

func main() {

	// run ever x seconds
	ticker := time.NewTicker(25 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Printf("%v", "\n -------------- RUNNING IN API -------------- \n")
		}
	}

}
