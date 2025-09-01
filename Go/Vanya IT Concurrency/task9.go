package main

import (
	"context"
	"fmt"
	"reflect"
)

// TASK 9
// Реализовать функцию orDone, которая направляет данные из канала in в возвращаемый канал out, пока канал in открыт
// и контекст не отменен

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
                // LEARN WHY WE NEED NESTED SELECT (TO RESOLVE out <- v blocking operation and cancel() case)
				select {
				case <-ctx.Done():
				case out <- v:
				}
			}
		}
	}()

	return out

}

func main() {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []interface{}
	for v := range orDone(context.Background(), ch) {
		res = append(res, v)
	}

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}

	for _, r := range res {
		fmt.Println("the value is", r)
	}

}
