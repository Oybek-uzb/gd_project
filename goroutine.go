package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go greeting("any word", ch)
	// time.Sleep(3 * time.Second) // if the main goroutine ends, all others goroutines also will be ended.
	// fmt.Println(<-ch) // if we use channels the main goroutine waits for this channels being completed.

	for i := range ch {
		fmt.Println(i)
	}
	// if channel does not be closed, it will cause for deadlock error.
}

func greeting(word string, ch chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
		ch <- true
	}
	close(ch)
}
