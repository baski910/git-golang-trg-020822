package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func ops1(ctx context.Context) error {
	fmt.Println(ctx)
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func ops2(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := ops1(ctx)
		if err != nil {
			cancel()
		}
	}()

	ops2(ctx)
}
