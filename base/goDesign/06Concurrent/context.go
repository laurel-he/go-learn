package main

import (
	"context"
	"fmt"
	"time"
)

func handleReq(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("__done__")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("__time after__")
	}
}

func testAllContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go handleReq(ctx)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("_____testAllContext___done__")
	}
}

func main() {
	testAllContext()
}
