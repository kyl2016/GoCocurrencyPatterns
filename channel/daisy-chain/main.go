package main

import "fmt"

func f(left, right chan int) {
	r := 1 + <-right
	println(r)
	left <- r
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)

	fmt.Println(<-leftmost)
}
