package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 1 OMIT
func race(name string, raceTo int) chan struct{} {
	done := make(chan struct{})

	go func() {

		// returns a channel will send the current time at a random interval
		t := time.Tick(time.Duration(int(rand.Float64() * float64(time.Second))))
		for i := 0; i < raceTo; i++ {
			<-t
			fmt.Printf("%s says: %d\n", name, i)
		}

		// signal that we are done counting
		done <- struct{}{}
		close(done)
	}()
	return done
}

// 2 OMIT

func main() {
	// START OMIT
	scott := race("Scott", 10)
	nolen := race("Nolen", 10)

	select {
	case <-scott:
		println("Scott wins, Nolen is a dumb dumb.")

	case <-nolen:
		println("Nolen wins, Scott is a dumb dumb.")
	}
	// END OMIT
}
