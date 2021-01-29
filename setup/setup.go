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
func InputNo() [3]int {
	var no [3]int

	for {
		var num int
		fmt.Println("enter 3 digit :")
		_, err := fmt.Scanf("%d\n", &num)
		if err != nil {
			fmt.Println("you have entered wrong")
			continue
		}

		idx := 0
		success := true
		// ex) [1st 123 % 10 = 3 , 123 / 10 = 12 ]	[2nd 12 % 10 = 2, 12 / 10 = 1] [3rd 1 % 10 = 1, 1/1=1 ]
		for num > 0 {
			n := num % 10
			num = num / 10

			//checking duplicated input number
			duplicated := false
			for i := 0; i < idx; i++ {
				if no[i] == n {
					duplicated = true
					break
				}
			}

			if duplicated {
				fmt.Println("cannot enter same digit")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("you have entered more than 3 number")
				success = false
				break
			}

			no[idx] = n
			idx++

		}
		if !success {
			continue
		}
		break

	}
	// to print out in order of input order
	no[0], no[2] = no[2], no[0]

	fmt.Println(no)
	return no

}

//CompareNo of random number and user input number
func CompareNo(no, inputNo [3]int) Result {
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if no[i] == inputNo[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	//fmt.Println(strikes)
	//fmt.Println(balls)
	return Result{strikes, balls}
}

//ShowResult of the game
func ShowResult(result Result) {
	fmt.Printf("%dS %dB\n", result.strikes, result.balls)
}

//EndGame shows the turn it took to finish
func EndGame(result Result) bool {
	return result.strikes == 3
}
