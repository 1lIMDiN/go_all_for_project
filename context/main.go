package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(2)
	go dots(ctx)
	wg.Wait()
}

func dots(ctx context.Context) {
	defer wg.Done()

	ctxAsterisk, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	go Asterisk(ctxAsterisk)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print(".")
			time.Sleep(500 *time.Millisecond)
		}
	}
}

func Asterisk(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Print("*")
			time.Sleep(500 *time.Millisecond)
		}
	}
}
