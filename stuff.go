package main

import (
	"context"
	"fmt"
	"time"

	"./api"
)

func doStuff(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := api.PerformOperation(ctx, 10, 20)
	if err != nil {
		return
	}

	fmt.Println(result)
	return
}
