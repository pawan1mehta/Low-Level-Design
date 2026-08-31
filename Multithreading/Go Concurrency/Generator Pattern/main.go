package main

import "fmt"

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
	nums := []int{1, 2, 3, 4, 5}

	data := generateData(nums)

	for data := range data {
		fmt.Println("data: ", data)
	}
}