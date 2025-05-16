package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}

		defer close(out)
	}()

	return out
}

func main() {
	out := randomGenerator(100)

	for w := range out {
		fmt.Println(w)
	}
}
