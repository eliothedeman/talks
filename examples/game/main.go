package main

import (
	"fmt"
	"math/rand"
)

type guess struct {
	guesser string
	guess   int
}
// 0 OMIT
func startGuessing(player string, max int) chan guess {
	out := make(chan guess)
	go func() {
		for {
			out <- guess{
				guess:   rand.Intn(max),
				guesser: player,
			}
		}
	}()

	return out
}
// 1 OMIT


func playGame(players []string) chan string {
	winner := make(chan string)
	go func() {
		round := 1
		scores := map[string]int{}
			//2 OMIT
		for {
			secret := rand.Intn(round * 100000)
			fanIn := make(chan guess)
			for i := range players {
				go func(i int) {
					out := startGuessing(players[i], round*100000)
					for g := range out {
						fanIn <- g
					}
				}(i)
			}
			for {
				g := <-fanIn
				if g.guess == secret {
					fmt.Printf("%s won round %d: correct guess %d\n", g.guesser, round, secret) // OMIT
					scores[g.guesser] += round
					for k, v := range scores { // OMIT
						fmt.Printf("%s: %d\n", k, v) //OMIT
					}// OMIT
					if scores[g.guesser] >= 100 {
						println("====SCORES====") // OMIT
						for k, v := range scores { // OMIT
							fmt.Printf("%s: %d\n", k, v) // OMIT
						} // OMIT
						winner <- g.guesser
						return
					}
					round++
					break
			//3 OMIT
				}
			}
		}

	}()
	return winner
}

func main() {

	// START OMIT
	winner := playGame([]string{
		"eliot",
		"nolen",
	})

	fmt.Printf("%s WINS!!!!\n", <-winner)
	// END OMIT
}
