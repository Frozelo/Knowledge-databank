package main

// TASK №16 NOT DONE YET
// Напишите функцию inc, которая принимает на вход канал, читает из него значения
// и записывает эти значения, увеличенные на единицу, в возвращаемый канал.
// Дополните функцию main созданием цепочки каналов, используя inc, так,
// чтобы программа завершалась без паники.

func main() {
	first := make(chan int)
	last := make(<-chan int)
	n := 10

    last = inc(first)
	first <- 0
	close(first)

    

	if n != <-last {
		panic("wrong code")
	}
}

// inc принимает канал для чтения и возвращает канал,
// в который будут отправляться считанные значения, увеличенные на 1.
func inc(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for v := range in {
            out <- v + 1
        }
    }()

	return out
}
