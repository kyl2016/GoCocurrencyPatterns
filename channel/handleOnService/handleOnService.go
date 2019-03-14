package main

import (
	"fmt"
)

// Our Boring function returns a channel that lets us communication with the Boring service it provides.
func main() {
	joe := Boring("Joe")
	ann := Boring("Ann")

	// Because of the synchronization nature of the channels, the two guys are taking turns,
	// not only in printing the values out, but also in executing them.
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You're both Boring; I'm leaving.")
}
