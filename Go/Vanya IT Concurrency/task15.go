package main

import (
	"fmt"
	"reflect"
)

// TASK 15
// Реализуйте методы структуры ringBuffer так, чтобы main отрабатывал без паники.

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{c: make(chan int, size)}
}

func (b *ringBuffer) write(v int) {
	select {
	case b.c <- v:
	default:
		<-b.c
		b.c <- v
	}
}

func (b *ringBuffer) close() {
	defer close(b.c)
}

func (b *ringBuffer) read() (v int, ok bool) {
	v, ok = <-b.c
	return
}

func main() {
	buff := newRingBuffer(3)
	for i := 1; i <= 6; i++ {
		buff.write(i)
	}
	buff.close()

	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}
	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
