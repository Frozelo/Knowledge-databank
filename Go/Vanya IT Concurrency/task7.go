package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TASK 7
// Напишите функцию getFirstResult, которая принимает контекст и запускает конкурентный поиск, возвращая
// первый доступный результат из replicas. Возвращать ошибку контекста, если контекст завершился раньше, чем стал
// доступен какой-то результат из replicas

// Напишите функцию getResults, которая запускает конкурентный поиск для каждого набора реплик из replicaKind,
// использую getFirstResults, и возвращает результат для каждого набора реплик

// THE FIRST WINS PATTERN
// The main idea of this task get the first result from the searching pairs. And transform this result into channel with getResults

type result struct {
	msg string
	err error
}

type search func() *result
type replicas []search

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	if len(replicas) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan *result)
	// searching in pairs
	for _, r := range replicas {
		go func(r search) {
			select {
			case <-ctx.Done():
			case ch <- r():
			}
		}(r)
	}

	select {
	case <-ctx.Done():
		return &result{err: ctx.Err()}
	case res := <-ch:
		return res
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	ch := make(chan *result)
	wg := sync.WaitGroup{}

	for _, replicas := range replicaKinds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- getFirstResult(ctx, replicas)
		}()
	}

	// waiting until done
	go func() {
		defer close(ch)
		wg.Wait()

	}()

	res := make([]*result, 0, len(replicaKinds))
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				return res
			}
			res = append(res, r)
		}
	}

	return res
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}

	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
