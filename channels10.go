package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("****************")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("cancel: ", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("****************")

	cancel()
	fmt.Println("## cancel() called ##")

	fmt.Println("Context: ", ctx)
	fmt.Println("context err: ", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("cancel: ", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("****************")
	fmt.Println("Notice how the context err is no longer nil after cancel() is called but now is context canceled")
}

//context can help you prevent leaking goroutines, with context you cancel one goroutine the other go routines end too.
