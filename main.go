package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/BrianH29/baseballGame/setup"
)

func main() {
	//to change random number, when the game is executed
	rand.Seed(time.Now().UnixNano())
	randNo := setup.RandomNo()

	count := 0
	for {
		count++
		inputNo := setup.InputNo()

		result := setup.CompareNo(randNo, inputNo)

		setup.ShowResult(result)

		if setup.EndGame(result) {
			fmt.Printf("it took %d turns to finish\n", count)
			break
		}

	}

}
