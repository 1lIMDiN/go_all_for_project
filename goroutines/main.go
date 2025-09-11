package main

import (
	"fmt"
)

func main() {
	chIn := make(chan int)
	chOut := make(chan int)

	go do(chIn, chOut)

	go func() {
		defer close(chIn)
		for i := 0; i <= 50; i++ {
			chIn <- i
		}
	}()

	for {
		v, ok := <-chOut
		if !ok {
			break
		}
		fmt.Print(v, " ")
	}
}

func do(in, out chan int) {
	defer close(out)
	for {
		v, ok := <-in
		if !ok {
			break
		}
		out <- v
	}
}
