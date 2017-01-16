package api

import (
	"context"
)

func PerformOperation(ctx context.Context, x, y int) (result int, err error) {
	performance := startPerforming(x, y)

	select {
	case result, _ = <-performance.resultChannel:
		return

	case <-ctx.Done(): // Timed out or cancelled?
		err = ctx.Err()
		performance.abort()
		return
	}
}
