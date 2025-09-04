package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

// TASK 12
// Реализовать интефейс waiter чтобы можно было конкурентно запскать maxParallels функций
// Доп от себя добавил ограничение на вызов run() после wait()

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	sem    chan struct{}
	mu     sync.Mutex
	wg     sync.WaitGroup
	closed atomic.Bool
	errs   []error
}

func (g *waitGroup) wait() error {
	g.closed.CompareAndSwap(false, true)

	g.wg.Wait()
	g.mu.Lock()

	defer g.mu.Unlock()

	if len(g.errs) == 0 {
		return nil
	}

	return errors.Join(g.errs...)
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	if g.closed.Load() {
		panic("waitGroup: run called after wait()")
	}

	select {
	case <-ctx.Done():
	case g.sem <- struct{}{}:
		if g.closed.Load() {
			<-g.sem
			panic("waitGroup: run called after wait()")
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		defer func() { <-g.sem }()

		if err := fn(ctx); err != nil {
			g.mu.Lock()
			g.errs = append(g.errs, err)
			g.mu.Unlock()
		}
	}()
}

func newGroupWait(maxParallel int) waiter {
	if maxParallel < 1 {
		maxParallel = 1
	}

	g := &waitGroup{sem: make(chan struct{}, maxParallel), wg: sync.WaitGroup{}, errs: make([]error, 0, maxParallel)}

	return g
}

func main() {
	g := newGroupWait(2)

	ctx := context.Background()
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")
	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})

	err := g.wait()
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wrong code")
	}


	fmt.Println(err)
}
