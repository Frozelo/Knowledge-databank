package main

import (
	"errors"
	"fmt"
	"sync"
)

// TASK 14
// Напишите функцию Run, которая запускает конкурентное выполнение функций fs и дожидается их окончания.
// Если одна или несколько функций из fs завершились с ошибкой, Run возвращает любую из них.

type fn func() error

func Run(fs ...fn) error {
	errorCh := make(chan error, 1)

	var wg sync.WaitGroup
	for _, f := range fs {
		wg.Add(1)
		go func(f fn) {
			defer wg.Done()
			if err := f(); err != nil {
				select {
				case errorCh <- err:
				default:
				}
			}
		}(f)
	}

	wg.Wait()
	close(errorCh)

	return <-errorCh
}

func main() {
	var err error
	expErr := errors.New("error")
	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err = Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}

	fmt.Println(err)
}
