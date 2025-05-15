package main

import "fmt"

func generator(in ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, w := range in {
			out <- w
		}
	}()

	return out
}

func main() {
	out := generator(1, 2, 3, 4, 5)

	for w := range out {
		fmt.Println(w)
	}

}
