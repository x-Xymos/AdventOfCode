package main

import (
	"fmt"
	"sync"
	"time"
)

type Results struct {
	c      chan int
	closed bool
	once   sync.Once
}

func (result *Results) Close() {
	result.once.Do(func() {
		close(result.c)
		result.closed = true
	})
}

func calculateInput(noun int, verb int, input []int, results *Results, wg *sync.WaitGroup) {
	defer wg.Done()
	input[1] = noun
	input[2] = verb
	for i := 0; i < len(input); i += 4 {
		if results.closed {
			return
		}
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case 99:
			if input[0] == 19690720 {
				results.c <- noun
				results.c <- verb
				results.c <- 0
				results.Close()
			}
			return
		default:
		}
	}
}

func main() {

	var input = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 10, 23, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 35, 9, 39, 2, 39, 13, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 9, 59, 63, 2, 6, 63, 67, 1, 13, 67, 71, 1, 9, 71, 75, 2, 13, 75, 79, 1, 79, 10, 83, 2, 83, 9, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 99, 5, 103, 1, 103, 2, 107, 1, 107, 10, 0, 99, 2, 0, 14, 0}

	results := Results{c: make(chan int)}

	var wg sync.WaitGroup

	start := make(chan struct{})
	for i := 0; i < len(input); i++ {
		for ii := 0; ii < len(input); ii++ {
			go func(index1 int, index2 int) {
				<-start // wait for the start channel to be closed
				wg.Add(1)
				calculateInput(index1, index2, append([]int(nil), input...), &results, &wg)
			}(i, ii)
		}
	} // at this point, all goroutines are ready to go - we just need to
	// tell them to start by closing the start channel
	close(start)

	go func() {
		for {
			res := <-results.c
			if res == 0 {
				fmt.Println("Finished")
				break
			}
			fmt.Println(res)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	wg.Wait()
}
