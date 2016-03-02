package main

import (
	"fmt"
	"sync"
)






func Square(n int) int {
	return n * n
}

type Maper func(int) int

func Map(f Maper, in chan int, p int) chan int {
	out := make(chan int)

	go func() {
		var wait sync.WaitGroup

		for i := 0; i < p; i++ {
			wait.Add(1)
			go func(count int) {
				fmt.Printf("Mapper %d started", count)
				for val := range in {
					out <- f(val)
				}
				wait.Done()
			}(i)
		}
		wait.Wait()
		close(out)
	}()

	return out
}

type Reducer func(chan int) chan int

func Sum(in chan int) chan int {
	out := make(chan int)
	go func() {
		total := 0
		for val := range in {
			total += val
		}
		out <- total
		close(out)
	}()
	return out
}

func Reduce(f Reducer, i chan int) chan int {
	return f(i)
}

func main() {
	vals := []int{
		0, 1, 5, 44, 99, 1, 44,
	}

	in := make(chan int)

	go func() {
		for _, v := range vals {
			in <- v
		}
		close(in)

	}()
	mapped := Map(Square, in, 1)
	reduced := Reduce(Sum, mapped)
	println(<-reduced)

}
