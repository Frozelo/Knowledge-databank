package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// TASK 8
// Написать функцию обертку для executeTask, которая:
// исполняет executeTask
// принимает аргументом контекст
// завершается либо в результате исполнения, либо в результате отмены контекста. В последнем случае вернуть ошибку контекста

const timeout = 100 * time.Millisecond

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}

func executeTaskWithTimeout(ctx context.Context) error {
	// сделать что-то чтобы понять, что все ок и executeTask выполнилась
	// <- struct{}

	ch := make(chan struct{}, 1)
	go func() {
        defer close(ch)
		executeTask()
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		return nil
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}
