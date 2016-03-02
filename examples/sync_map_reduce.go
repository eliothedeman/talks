package main

func Square(n int) int {
	return n * n
}

type Maper func(int) int

func Map(f Maper, i []int) []int {
	for x, y := range i {
		i[x] = f(y)
	}
	return i
}

type Reducer func([]int) int

func Sum(i []int) int {
	total := 0
	for _, x := range i {
		total += x
	}

	return total
}

func Reduce(f Reducer, i []int) int {
	return f(i)
}

func main() {
	vals := []int{
		0, 1, 5, 44, 99, 1, 44,
	}

	println(Reduce(Sum, Map(Square, vals)))

}
