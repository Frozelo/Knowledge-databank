package main

import (
	"fmt"
	"log"
	"sync"
)

// TASK 2.1.1
// Напишите функцию merge и fillChan

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))

	for i, ch := range cs {
		log.Println("startting goroutine with id", i)
		go func(i int, ch <-chan int) {
			log.Printf("goroutine %v started \n", i)
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(i, ch)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func fillChan(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range 10 {
			ch <- v
		}
	}()

	return ch
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)
	for v := range d {
		fmt.Println(v)
	}

}
