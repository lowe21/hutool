package device

import (
	"context"
	"sync"

	"hutool/internal/pkg/device/detector"
)

var (
	device *detector.Detector
	once   sync.Once
)

func instance() *detector.Detector {
	once.Do(func() {
		device = detector.New(&detector.Config{})
	})

	return device
}

// Listener 监听器
func Listener(ctx context.Context, datasetFunc func(context.Context, [][]float64), errFunc func(context.Context, error)) {
	instance().Listener(ctx, datasetFunc, errFunc)
}

// StandbyData 待机数据
func StandbyData() []float64 {
	return instance().StandbyData()
}
