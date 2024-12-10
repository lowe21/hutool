package detector

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/util/gconv"

	detector "hutool/internal/pkg/device"
	"hutool/internal/pkg/websocket"
	"hutool/internal/service"
)

func init() {
	service.RegisterDetector(&detectorLogic{})
}

type detectorLogic struct{}

// AnalyzeData 分析数据
func (detector *detectorLogic) AnalyzeData(_ context.Context, dataset [][]float64) {
	if len(dataset) == 0 {
		return
	}

	// 计算数据
	data := make([]float64, len(dataset[0]))
	for _, item := range dataset {
		for index, value := range item {
			data[index] += value
		}
	}
	for index := range data {
		data[index] /= gconv.Float64(len(dataset))
	}

	type Data struct {
		Type string  `json:"type"`
		X    int     `json:"x"`
		Y    float64 `json:"y"`
	}

	newData := make([]*Data, 0, (len(data)+2)*2)
	x := 0

	// 数据头
	newData = append(newData, &Data{
		Type: "benchmark",
		X:    x,
	})
	newData = append(newData, &Data{
		Type: "testing",
		X:    x,
	})

	switch {
	// 是否为待机模式
	case detector.isStandbyMode(data):
		for range data {
			x++
			newData = append(newData, &Data{
				Type: "benchmark",
				X:    x,
			})
			newData = append(newData, &Data{
				Type: "testing",
				X:    x,
			})
			x++
		}
	default:
		// TODO
		for _, value := range data {
			x++
			newData = append(newData, &Data{
				Type: "benchmark",
				X:    x,
			})
			newData = append(newData, &Data{
				Type: "testing",
				X:    x,
				Y:    value,
			})
			x++
		}
	}

	// 数据尾
	newData = append(newData, &Data{
		Type: "benchmark",
		X:    x,
	})
	newData = append(newData, &Data{
		Type: "testing",
		X:    x,
	})

	// 通知数据
	websocket.Notice(websocket.Message("data", "newData"))
}

// Error 错误
func (*detectorLogic) Error(_ context.Context, err error) {
	// 通知错误
	websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
}

// isStandbyMode 是否为待机模式
func (*detectorLogic) isStandbyMode(data []float64) bool {
	// 待机数据
	standbyData := detector.StandbyData()
	if len(standbyData) != len(data) {
		return false
	}

	for index, value := range data {
		if value != standbyData[index] {
			// 误差值
			if math.Abs(value-standbyData[index]) > 0.1 {
				return false
			}
		}
	}

	return true
}
