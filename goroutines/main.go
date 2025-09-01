package main

import (
	"fmt"
	"sync"
	"time"
)

const Count = 4

var wg sync.WaitGroup

func main() {
	ch := make(chan int, Count)

	for i := 0; i < Count; i++ {
		wg.Add(1)

		go worker(ch, i)
	}

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}

func worker(ch <-chan int, i int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("ch =", i, "v = ", v)
		time.Sleep(50 * time.Millisecond)
	}
}
