package main

import (
	"runtime"
)

func doSomething() {

}

func main() {
	// START OMIT
	go doSomething()
	// END OMIT
	println(runtime.NumGoroutine())
}
