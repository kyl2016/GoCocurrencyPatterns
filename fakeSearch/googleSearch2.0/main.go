package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := Google("happy")
	fmt.Printf("%+v", r)
}

var (
	Web   = fakeSearch("web")
	Video = fakeSearch("video")
	Image = fakeSearch("image")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return results
}
