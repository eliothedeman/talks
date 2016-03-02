package main

import (
	"fmt"
	"strings"
)

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

func main() {

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
}
