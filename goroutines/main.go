package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j <= 10; j++ {
				m.Store(j, j*j)
			}
		}()
	}
	wg.Wait()

	m.Range(func(key, value any) bool {
		if v, ok := value.(int); ok {
			fmt.Println(v)
		}
		fmt.Printf("%v x %v = %v\n", key, key, value)
		return true
	})
}
