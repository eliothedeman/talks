package main

import (
	"fmt"
	"math/rand"
)

type guess struct {
	guesser string
	guess   int
}

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

func playGame(p []string) chan string {
	winner := make(chan string)
	go func() {
		round := 1
		scores := map[string]int{}
		for {
			secret := rand.Intn(round * 100000)
			fanIn := make(chan guess)

			for i := range p {
				go func(i int) {
					out := startGuessing(p[i], round*100000)
					for g := range out {
						fanIn <- g
					}
				}(i)
			}

			for {
				g := <-fanIn
				if g.guess == secret {
					fmt.Printf("%s won round %d: correct guess %d\n", g.guesser, round, secret)
					scores[g.guesser] += round
					for k, v := range scores {
						fmt.Printf("%s: %d\n", k, v)
					}
					if scores[g.guesser] >= 100 {
						println("====SCORES====")
						for k, v := range scores {
							fmt.Printf("%s: %d\n", k, v)
						}
						winner <- g.guesser
						return
					}
					round++
					break

				}
			}
		}

	}()
	return winner
}

func main() {
	winner := playGame([]string{
		"eliot",
		"nolen",
	})

	fmt.Printf("%s WINS!!!!\n", <-winner)
}
