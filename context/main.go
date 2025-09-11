package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main () {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	wg.Add(2)
	go dots(ctx)
	wg.Wait()
}

func dots(ctx context.Context) {
	defer wg.Done()
	ctxAster, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()
	go Aster(ctxAster)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Print(".")
		}
	}
}

func Aster(ctx context.Context) {
	defer wg.Done()
	for {
		select{
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Print("*")
		}
	}
}