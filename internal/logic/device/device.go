package device

import (
	"context"
	"math"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"

	"hutool/internal/pkg/websocket"
	"hutool/internal/service"
	"hutool/internal/util"
)

func init() {
	service.RegisterDevice(&deviceLogic{})
}

type deviceLogic struct{}

// Listener 设备监听器
func (device *deviceLogic) Listener(ctx context.Context) (err error) {
	defer func() {
		// 异常处理
		if exception := recover(); exception != nil {
			err = util.Error(exception)
		}

		// 错误处理
		if err != nil {
			websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
		}

		// 重启监听器
		gtimer.SetTimeout(context.TODO(), device.restartInterval(ctx), func(context.Context) {
			err = service.Device().Listener(ctx)
		})
	}()

	// 获取端口列表
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return
	}

	// 设备端口名称
	var portName string

	for _, port := range ports {
		// 匹配设备VID和PID
		if port.VID == device.vid(ctx) && port.PID == device.pid(ctx) {
			portName = port.Name
			break
		}
	}

	if portName == "" {
		return util.Error("Serial port not found")
	}

	// 打开设备端口
	port, err := serial.Open(portName, &serial.Mode{
		BaudRate: device.baudRate(ctx),
	})
	if err != nil {
		return
	}

	// 数据集
	dataset := make([][]float64, 0, gconv.Int(math.Ceil(gconv.Float64(time.Second.Milliseconds())/gconv.Float64(device.dataInterval(ctx)))))

	for {
		// 缓冲区
		size := 64
		buffer := make([]byte, size)

		// 读取数据
		size, err = port.Read(buffer)
		if err != nil {
			break
		}

		// 解析数据
		data := service.Device().ParseData(ctx, buffer[:size])
		if len(data) > 0 {
			dataset = append(dataset, data)
		}

		// 重置数据集
		if len(dataset) == cap(dataset) {
			websocket.Notice(websocket.Message("data", service.Device().BuildData(ctx, dataset)))
			dataset = dataset[:0]
		}
	}

	return
}

// ParseData 解析数据
func (device *deviceLogic) ParseData(ctx context.Context, buffer []byte) (data []float64) {
	// 数据大小
	dataSize := device.dataSize(ctx)
	if dataSize != len(buffer) {
		return
	}

	// 数据帧头
	dataHead := garray.NewFrom(gconv.Interfaces(device.dataHead(ctx)))
	bufferHead := garray.NewFrom(gconv.Interfaces(buffer[:dataHead.Len()]))
	if dataHead.Join(",") != bufferHead.Join(",") {
		return
	}

	// 数据帧尾
	dataTail := garray.NewFrom(gconv.Interfaces(device.dataTail(ctx)))
	bufferTail := garray.NewFrom(gconv.Interfaces(buffer[dataSize-dataTail.Len():]))
	if dataTail.Join(",") != bufferTail.Join(",") {
		return
	}

	data = make([]float64, 0, dataSize-dataHead.Len()-dataTail.Len())

	for _, value := range buffer[dataHead.Len() : dataSize-dataTail.Len()] {
		data = append(data, gconv.Float64(value)/16)
	}

	return
}

// BuildData 构建数据
func (*deviceLogic) BuildData(ctx context.Context, dataset [][]float64) any {
	if len(dataset) == 0 {
		return nil
	}

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
		X int     `json:"x"`
		Y float64 `json:"y"`
	}

	dataList := make([]*Data, 0, 3*len(data))
	x := 1

	for range data {
		dataList = append(dataList, &Data{
			X: x,
		})
		x++
	}

	switch {
	case service.Device().IsStandby(ctx, data):
		for range data {
			dataList = append(dataList, &Data{
				X: x,
			})
			x++
		}
	default:
		for _, value := range data {
			dataList = append(dataList, &Data{
				X: x,
				Y: value,
			})
			x++
		}
	}

	for range data {
		dataList = append(dataList, &Data{
			X: x,
		})
		x++
	}

	return dataList
}

// IsStandby 是否为待机
func (device *deviceLogic) IsStandby(ctx context.Context, data []float64) bool {
	standbyData := device.standbyData(ctx)
	if len(data) != len(standbyData) {
		return false
	}

	for index, value := range data {
		if value != standbyData[index] {
			if math.Abs(value-standbyData[index]) > 0.1 {
				return false
			}
		}
	}

	return true
}

// vid 供应商ID
func (*deviceLogic) vid(ctx context.Context) string {
	return g.Config().MustGet(ctx, "device.vid", "1A86").String()
}

// pid 产品ID
func (*deviceLogic) pid(ctx context.Context) string {
	return g.Config().MustGet(ctx, "device.pid", "7523").String()
}

// baudRate 波特率
func (*deviceLogic) baudRate(ctx context.Context) int {
	return g.Config().MustGet(ctx, "device.baudRate", 115200).Int()
}

// dataInterval 数据间隔
func (*deviceLogic) dataInterval(ctx context.Context) int {
	return g.Config().MustGet(ctx, "device.dataInterval", 100).Int()
}

// dataSize 数据大小
func (*deviceLogic) dataSize(ctx context.Context) int {
	return g.Config().MustGet(ctx, "device.dataSize", 14).Int()
}

// dataHead 数据帧头
func (*deviceLogic) dataHead(ctx context.Context) []byte {
	return g.Config().MustGet(ctx, "device.dataHead", []byte{255, 129}).Bytes()
}

// dataHead 数据帧尾
func (*deviceLogic) dataTail(ctx context.Context) []byte {
	return g.Config().MustGet(ctx, "device.dataTail", []byte{204, 90}).Bytes()
}

// restartInterval 重启间隔
func (*deviceLogic) restartInterval(ctx context.Context) time.Duration {
	return g.Config().MustGet(ctx, "device.restartInterval", "5s").Duration()
}

// standbyData 待机数据
func (*deviceLogic) standbyData(ctx context.Context) []float64 {
	return g.Config().MustGet(ctx, "device.standbyData", []float64{10.5, 7.8125, 7.8125, 7.1875, 7.875, 7.625, 7.875, 7.875, 7.875, 10.125}).Float64s()
}
