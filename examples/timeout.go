package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// START OMIT
	wait := func() <-chan time.Time {
		return time.After(time.Duration(rand.Float64() * 2 * float64(time.Second)))
	}

	select {
	case <-wait():
		println("Success.")

	case <-time.After(time.Second):
		println("FAILURE: Timeout hit")
	}
	// END OMIT
}
