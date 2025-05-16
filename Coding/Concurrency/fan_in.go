package main

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...chan int) chan int {
	var wg sync.WaitGroup

	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch chan int) {
			defer wg.Done()

			for id := range ch {
				out <- id
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	out := joinChannels(a, b, c)

	for w := range out {
		fmt.Println(w)
	}

}
