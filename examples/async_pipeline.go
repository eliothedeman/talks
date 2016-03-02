package main

import (
	"fmt"
	"strings"
)

// 1 OMIT
func Quote(msg chan string) chan string {
	out := make(chan string)
	go func() {
		for s := range msg {
			out <- fmt.Sprintf("'%s'", s)
		}
		close(out)
	}()
	return out
}

// 2 OMIT

// 3 OMIT
func Cap(msg chan string) chan string {
	out := make(chan string)
	go func() {
		for s := range msg {
			out <- strings.ToUpper(s)
		}
		close(out)
	}()
	return out
}

// 4 OMIT

// 5 OMIT
func Explode(msg chan string) chan string {
	out := make(chan string)
	go func() {
		for s := range msg {
			e := make([]byte, (len(s)*2)-1)
			for i := 0; i < len(s); i++ {
				e[i*2] = s[i]
				if i != len(s)-1 {
					e[(i*2)+1] = ' '
				}
			}
			out <- string(e)
		}
		close(out)
	}()
	return out
}

// 6 OMIT

func main() {

	// START OMIT
	strs := []string{
		"happy",
		"birthday",
		"to",
		"you",
	}
	in := make(chan string)

	go func() {
		for _, s := range strs {
			in <- s
		}
		close(in)
	}()

	for s := range Quote(Explode(Cap(in))) {
		println(s)
	}
	// END OMIT
}
