package main

import (
	"fmt"
	"strings"
)

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

	// START OMIT
	strs := []string{
		"happy",
		"birthday",
		"to",
		"you",
	}

	for _, s := range strs {
		println(Quote(Explode(Cap(s))))
	}
	// END OMIT
}
