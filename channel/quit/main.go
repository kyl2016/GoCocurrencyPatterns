package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)

	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring(msg string, quit <-chan bool) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			select {
			case ch <- fmt.Sprintf("%s: %d", msg, i):
			// do nothing
			case <-quit:
				return
			}
		}
	}()

	return ch
}
