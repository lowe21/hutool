// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDevice interface {
		// Listener 设备监听器
		Listener(ctx context.Context) (err error)
		// ParseData 解析数据
		ParseData(ctx context.Context, buffer []byte) (data []float64)
		// BuildData 构建数据
		BuildData(ctx context.Context, dataset [][]float64) any
		// IsStandby 是否为待机
		IsStandby(ctx context.Context, data []float64) bool
	}
)

var (
	localDevice IDevice
)

func Device() IDevice {
	if localDevice == nil {
		panic("implement not found for interface IDevice, forgot register?")
	}
	return localDevice
}

func RegisterDevice(i IDevice) {
	localDevice = i
}
