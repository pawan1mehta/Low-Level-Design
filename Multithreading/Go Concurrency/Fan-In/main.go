package main

import (
	"fmt"
	"sync"
)

func merge(ch ...<-chan int) <-chan int {
	outChan := make(chan int)
	var wg sync.WaitGroup

	for _, c := range ch {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for num := range c {
				outChan <- num
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}

func generateData(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()
	return ch
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}

	data1 := generateData(nums1)
	data2 := generateData(nums2)

	for num := range merge(data1, data2) {
		fmt.Println(num)
	}
}
