package api

import (
	"time"
)

type performance struct {
	resultChannel <-chan int
}

func startPerforming(x, y int) *performance {
	c := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		c <- x + y
	}()

	return &performance{
		resultChannel: c,
	}
}

func (*performance) abort() {}
