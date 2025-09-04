package main

import (
	"context"
	"fmt"
	"sync"
)

func fanIn(ctx context.Context, chs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}
		}(ch)

	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	out := joinChannels(ctx, a, b, c)

	for w := range out {
		fmt.Println(w)
	}

}
