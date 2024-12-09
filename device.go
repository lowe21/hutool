package main

import (
	"context"

	"hutool/internal/service"
)

type Device struct{}

func (*Device) listener(ctx context.Context) {
	_ = service.Device().Listener(ctx)
}
