package main

import (
	"context"

	detector "hutool/internal/pkg/device"
	"hutool/internal/service"
)

type Device struct{}

func (*Device) listener(ctx context.Context) {
	detector.Listener(ctx, func(ctx context.Context, dataset [][]float64) {
		service.Detector().AnalyzeData(ctx, dataset)
	}, func(ctx context.Context, err error) {
		service.Detector().Error(ctx, err)
	})
}
