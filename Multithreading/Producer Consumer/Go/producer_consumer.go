package main

import (
	"fmt"
	"sync"
)

func main() {
	queue := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	// producer
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			queue <- i
			fmt.Println("produced: ", i)
		}
		close(queue)
	}()

	// consumer
	go func() {
		defer wg.Done()
		for item := range queue {
			fmt.Println("consumed: ", item)
		}
	}()

	wg.Wait()
}
