package setup

import (
	"fmt"
	"math/rand"
)

//Result type
type Result struct {
	strikes int
	balls   int
}

//RandomNo random number of 3 between 1~9
func RandomNo() [3]int {
	var no [3]int

	for i := 0; i < 3; i++ {
		rand := rand.Intn(10)
		for {
			//duplication check
			duplicated := false
			for j := 0; j < i; j++ {
				if no[j] == rand {
					duplicated = true
					break
				}
			}

			if !duplicated {
				no[i] = rand
				break
			}
		}
	}
	//after the test of the game works fine? errase println
	fmt.Println(no)
	return no
}

//InputNo user input 3 digit
func InputNo() {

}

//CompareNo of random number and user input number
func CompareNo(no, inputNo [3]int) bool {
	return true
}

//ShowResult of the game
func ShowResult(result Result) bool {
	return true
}

//EndGame shows the turn it took to finish
func EndGame(result Result) {

}
