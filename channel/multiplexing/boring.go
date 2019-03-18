package main

import (
	"math/rand"
	"strconv"
	"time"
)

func Boring(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- msg + strconv.Itoa(i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return ch
}
