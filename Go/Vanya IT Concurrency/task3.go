package main

import (
	"context"
	"fmt"
)

// TASK 3 Написать функцию generator и squarer, которые могут отменяться по контексту

func squarer(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case out <- v * v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func generator(ctx context.Context, in ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()

	return out
}

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3, 4, 5))
	for x := range pipeline {
		fmt.Println(x)
	}
}
