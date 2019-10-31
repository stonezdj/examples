package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1000 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
