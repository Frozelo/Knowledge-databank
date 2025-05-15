package main

import (
	"context"
	"fmt"
	"sync"
)

// WorkerFunc определяет подпись функции, выполняемой каждым воркером.
type WorkerFunc func(ctx context.Context, id, job int) (int, error)

// worker обрабатывает задачи из канала jobs и отправляет результаты в results.
func worker(ctx context.Context, id int, fn WorkerFunc, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Выполняем функцию и обрабатываем возможную ошибку
			res, err := fn(ctx, id, job)
			if err != nil {
				fmt.Printf("Worker %d error on job %d: %v\n", id, job, err)
				continue
			}
			results <- res
		}
	}
}

// dispatcher запускат пул воркеров и возвращает канал с результатами.
func dispatcher(ctx context.Context, workersCount, jobsCount int, fn WorkerFunc) <-chan int {
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)
	var wg sync.WaitGroup

	// Стартуем воркеров
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(ctx, i, fn, jobs, results, &wg)
	}

	// Производитель задач
	go func() {
		defer close(jobs)
		for j := 1; j <= jobsCount; j++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- j:
			}
		}
	}()

	// Закрываем results после завершения всех воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workersCount := 4
	jobsCount := 100
	
	// Функция возведения в квадрат с демонстрацией контекста
	fn := func(ctx context.Context, id, job int) (int, error) {
		// Здесь можно учитывать ctx для отмены внутри работы
		return job * job, nil
	}

	fmt.Printf("Запускаем dispatcher: %d воркеров, %d задач\n", workersCount, jobsCount)
	results := dispatcher(ctx, workersCount, jobsCount, fn)

	for res := range results {
		fmt.Println("Result:", res)
	}

	fmt.Println("Все задачи обработаны.")
}
