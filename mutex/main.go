package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mutex   sync.Mutex
	counter map[string]int
}

func (c *Counter) Inc(key string) {
	c.mutex.Lock()
	c.counter[key]++
	c.mutex.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter[key]
}

func main() {
	const key = "test"
	c := Counter{counter: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(c.Value(key))
}
