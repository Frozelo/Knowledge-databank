package main

import (
	"context"
	"fmt"
	"math/rand"
)

// TASK 4 Напишите функцию repeatFN и take
// Функция repeatFn бесконечно вызывает функцию fn и пишет результат ее работы в возвращаемый канал. Прекращает работу
// раньше, если контекст отменен

// Функция take читает не более num из канала in, пока in открыт, и пишет значение в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}()

	return out
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} { return rand.Int() }
	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, rand), 3) {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}

	fmt.Println("you doing great because the len res is", len(res))
}
