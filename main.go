package main

import (
	"context"
	"fubuki-go/api"
	"fubuki-go/log"
	"fubuki-go/manager"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	go func() {
		cancel(api.Run())
	}()

	go func() {
		cancel(manager.Manager.Run())
	}()

	<-ctx.Done()
	log.Fatal("goroutine was exited", zap.Error(context.Cause(ctx)))
}
