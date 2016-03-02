package main

import (
	"fmt"
	"strings"
)

// 1 OMIT
func Quote(msg string) string {
	return fmt.Sprintf("'%s'", msg)
}

// 2 OMIT

// 3 OMIT
func Cap(msg string) string {
	return strings.ToUpper(msg)
}

// 4 OMIT

// 5 OMIT
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

// 6 OMIT

func main() {

	// START OMIT
	println(Quote("hello"))
	println(Explode("hello"))
	println(Cap("hello"))

	// run them all together
	println(Quote(Explode(Cap("hello"))))
	// END OMIT

}
