package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 3
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	semaphore := make(chan struct{}, n)
	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)

		semaphore <- struct{}{} // block if n jobs is already running

		go func(job int) {
			defer wg.Done()
			defer func() { <-semaphore }()
			fmt.Println("run ", job)
		}(job)
	}

	wg.Wait()
}
