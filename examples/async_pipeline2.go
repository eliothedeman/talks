package main

import (
	"fmt"
	"strings"
)

type StrFunc func(string) string
type ConcurrentStrFunc func(chan string) chan string

func Wrap(f StrFunc) ConcurrentStrFunc {
	return func(msg chan string) chan string {
		out := make(chan string)
		go func() {
			for s := range msg {
				out <- f(s)
			}
			close(out)
		}()
		return out
	}
}

func Quote(msg string) string {
	return fmt.Sprintf("'%s'", msg)
}

func Cap(msg string) string {
	return strings.ToUpper(msg)
}

func Explode(msg string) string {

	e := make([]byte, (len(msg)*2)-1)
	for i := 0; i < len(msg); i++ {
		e[i*2] = msg[i]
		if i != len(msg)-1 {
			e[(i*2)+1] = ' '

		}
	}

	return string(e)
}

func main() {
	strs := []string{
		"never",
		"going",
		"to",
		"give",
		"you",
		"up",
		"never",
		"going",
		"to",
		"let",
		"you",
		"down",
	}
	in := make(chan string)

	go func() {
		for _, s := range strs {
			in <- s
		}
		close(in)
	}()

	cQuote := Wrap(Quote)
	cExplode := Wrap(Explode)
	cCap := Wrap(Cap)

	for s := range cQuote(cExplode(cCap(in))) {
		println(s)
	}

}
