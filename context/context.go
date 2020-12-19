package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Context is feature to prevent to remain created goroutine when caller exit due to timeout.

	// get empty Context
	ctx := context.Background()
	// set timeout to Context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := hello(ctx); err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	fmt.Println("success!")
}

func hello(ctx context.Context) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan struct{}, 1)

	go func() {
		// sleep to confirm the behavior of Context Cancel
		//
		// If sleep time longer than the timeout period it initially set(e.g. 10 sec), ctx.Done() will be called.
		time.Sleep(1 * time.Second)
		fmt.Println("hello")
		ch <- struct{}{}
	}()

	select {
	// timeout
	case <-c.Done():
		return c.Err()
	case <-ch:
		return nil
	}
}
