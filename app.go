package main

import (
	"context"
	"os"
	"os/signal"
)

func app(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go waitForSignalAndInvoke(os.Interrupt, cancel)

	doStuff(ctx)
}

func waitForSignalAndInvoke(number os.Signal, cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, number)
	<-c
	cancel()
}
