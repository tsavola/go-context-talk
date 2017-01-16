package main

import (
	"context"
)

func main() {
	ctx := context.Background() // No timeout, cancellation, or values

	app(ctx)
}
