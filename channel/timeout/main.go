package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boring(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		sleepAmount := 0
		for i := 0; i < 20; i++ {
			ch <- msg + " " + strconv.Itoa(i)
			random := rand.Intn(1e3)
			sleepAmount += random
			fmt.Println(sleepAmount, time.Now())
			time.Sleep(time.Duration(random) * time.Millisecond)
			fmt.Println(time.Now())
		}
	}()

	return ch
}
