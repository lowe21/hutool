// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDetector interface {
		// AnalyzeData 分析数据
		AnalyzeData(_ context.Context, dataset [][]float64)
		// Error 错误
		Error(_ context.Context, err error)
	}
)

var (
	localDetector IDetector
)

func Detector() IDetector {
	if localDetector == nil {
		panic("implement not found for interface IDetector, forgot register?")
	}
	return localDetector
}

func RegisterDetector(i IDetector) {
	localDetector = i
}
