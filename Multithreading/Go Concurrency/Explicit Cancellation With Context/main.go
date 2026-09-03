package main

import (
	"context"
	"fmt"
	"time"
)

func processRequest(ctx context.Context, requestPayload string) (string, error) {
	processCh := make(chan string, 1)

	go func() {
		time.Sleep(300 * time.Millisecond)
		select {
		case processCh <- "Processed: " + requestPayload:

		case <-ctx.Done():
			return
		}
	}()

	select {
	case result := <-processCh:
		return result, nil
	case <-ctx.Done():
		return "nil", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	result, err := processRequest(ctx, "request payload")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Result: ", result)
}
