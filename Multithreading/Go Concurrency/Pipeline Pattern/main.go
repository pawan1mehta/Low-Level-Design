package main

import "fmt"

func generateData(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func processData(data <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for d := range data {
			out <- d
		}
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := generateData(nums)
	out := processData(ch)
	for v := range out {
		fmt.Println(v)
	}
}
