package main

import (
	"fmt"
	"math/rand"
	"time"
)

func race(name string, smart bool) chan int {
	out := make(chan int)
	var t <-chan time.Time
	if smart {
		t = time.Tick(time.Second / 2)
	} else {
		t = time.Tick(time.Duration(int(rand.Float64() * 1000000000)))
	}
	go func() {
		i := 0
		for {
			<-t
			out <- i
			fmt.Printf("%s: %d\n", name, i)
			i++
		}
	}()
	return out

}

func main() {
	eliot := race("Eliot", true)
	nolen := race("Nolen", false)

	// START OMIT
	count := 0
	for {
		select {
		case count = <-eliot:
			if count == 9 {
				println("Eliot wins")
			}

		case count = <-nolen:
			if count == 9 {
				println("Nolen wins")
			}
		}

		if count == 9 {
			break
		}
	}
	// END OMIT

}
