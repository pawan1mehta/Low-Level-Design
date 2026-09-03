package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	urls := []string{
		"https://go.dev",
		"https://golang.org",
		"https://invalid.example",
	}

	for _, url := range urls {
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return err
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Println(url, resp.Status)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("failed: ", err)
	}
}
