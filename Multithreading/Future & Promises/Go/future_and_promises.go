package main

import "fmt"

type Result struct {
	Value string
	Err   error
}

func ExecuteTask(data string) <-chan Result {
	ch := make(chan Result, 1)

	go func() {
		defer close(ch)
		ch <- Result{Value: "Processed: " + data, Err: nil}
	}()

	return ch
}

func main() {
	future := ExecuteTask("task1")

	fmt.Println("waiting for answer...")
	result := <-future
	fmt.Println("result: ", result.Value)
}
