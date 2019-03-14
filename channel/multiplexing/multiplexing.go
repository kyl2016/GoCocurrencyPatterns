package main

// Fan-in

func main() {
	c := fanIn(Boring("Joe"), Boring("Ann"))
	for i := 0; i < 10; i++ {
		println(<-c)
	}
	println("You're both Boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}
