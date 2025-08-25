package main

import (
	"fmt"
)

// TASK 5 Написать worker pool
// Нужно выполнить параллельно numJobs заданий, используя numWorkers горутин, которые запущены единожды за время
// выполнения программы

const numJobs = 5
const numWorkers = 3

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- f(j)
    }
}

func main() {
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    wg := sync.WaitGroup{}

    multiplier := func(x int) int {
        return x * 10
    }

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(){
            defer wg.Done()
            worker(multiplier, jobs, results)
        }()
    }

    go func() {
        defer close(jobs)
        for i := 0; i < numJobs; i++ {
            jobs <- i
        }
    }()

    go func() {
        wg.Wait()
        close(results)
    }()

    for v := range results {
        fmt.Println(v)
    }
}
