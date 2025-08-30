package main

import "fmt"

func generator(ch chan int, done chan struct{}) {
	val := 0
	dif := 0


	for {

		select {
		case ch <- val:

		case <-done:
			return
		}

		dif++
		val += dif

	}
}

func main() {
	ch := make(chan int, 15)
	done := make(chan struct{})

	go func() {
		defer close(done)

		for i := 0; i < 15; i++ {
			fmt.Print(<-ch, " ")
		}
	}()

	generator(ch, done)
}
