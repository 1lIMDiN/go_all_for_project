package main

import (
	"fmt"
	"sync"
	"time"
)

type Cashe struct {
	mu sync.Mutex
	m  map[int]int
}


func main() {
	cashe := Cashe{
		m: make(map[int]int),
	}

	for i := 0; i < 20; i++ {

		go func() {
			for j := 0; j < 1000; j++ {
				cashe.Get(j)
			}
		}()

	}

	time.Sleep(1 * time.Second)
	fmt.Println(cashe.m)
}

func (c *Cashe) Get(i int) int {

	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.m[i]
	if ok {
		return v
	}

	v = 2 * i
	c.m[i] = v
	return v
}
