package main

import "fmt"

func main() {
	loop(10, fibonacci())
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y //0, 1 -> 1, 2, 3, 5
		return x
	}
}

func loop(n int, f func() int) {
	if n > 0 {
		fmt.Println(f())
		loop(n-1, f)
	}
}
