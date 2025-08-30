package main

import (
	"fmt"
)

// TASK 7 NOT DONE YET!!!
// Напишите функцию getFirstResult, которая принимает контекст и запускает конкурентный поиск, возвращая
// первый доступный результат из replicas. Возвращать ошибку контекста, если контекст завершился раньше, чем стал
// доступен какой-то результат из replicas 

// Напишите функцию getResults, которая запускает конкурентный поиск для каждого набора реплик из replicaKind, 
// использую getFirstResults, и возвращает результат для каждого набора реплик

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
    // todo 
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
    ch := make(chan *result )
    wg := sync.WaitGroup{}

    for _, replica := range replicaKinds {
        wg.Add(1)
        go func(){
            defer wg.Done()
            ch <- getFirstResult(ctx, replicas)
        }()
    }

    wg.Wait()


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
}

func main() {
    ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
    replicaKinds := []replicas{
        replicas{fakeSearch("web1"), fakeSearch("web2")},
        replicas{fakeSearch("image1"), fakeSearch("image2")},
        replicas{fakeSearch("video1"), fakeSearch("video2")},
    }

    for _, res := range getResults(ctx, replicaKinds) {
        fmt.Println(res.Msg, res.err)
    }
}
