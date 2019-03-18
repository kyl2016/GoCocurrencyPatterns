package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)

	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"

	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			select {
			case ch <- fmt.Sprintf("%s: %d", msg, i):
			// do nothing
			case <-quit:
				// cleanup()
				quit <- "See you!"
				return
			}
		}
	}()

	return ch
}
