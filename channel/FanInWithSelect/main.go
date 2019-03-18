package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	waitForIt := make(chan bool)

	c := fanIn(Boring(Message{"Joe", waitForIt}), Boring(Message{"Ann", waitForIt}))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println("result 1:" + msg1.str)
		msg2 := <-c
		fmt.Println("result 2:" + msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

	println("You're both Boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			select {
			case r := <-input1:
				fmt.Printf("fanIn 1 %s\n", r.str)
				c <- r
			case r := <-input2:
				fmt.Printf("fanIn 2 %s\n", r.str)
				c <- r
			}
		}
	}()

	return c
}

type Message struct {
	str  string
	wait chan bool
}

func Boring(msg Message) <-chan Message {
	ch := make(chan Message)

	go func() {
		for i := 0; i < 5; i++ {
			waitForIt := make(chan bool)

			ch <- Message{fmt.Sprintf("create message %s: %d", msg.str, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return ch
}
