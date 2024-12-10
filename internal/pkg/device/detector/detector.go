package detector

import (
	"context"
	"math"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"

	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"

	"hutool/internal/util"
)

func New(config *Config) *Detector {
	// 初始化配置
	if config == nil {
		config = &Config{}
	}
	config.Init()

	return &Detector{
		config: config,
	}
}

type Detector struct {
	config *Config // 配置
}

// Listener 监听器
func (detector *Detector) Listener(ctx context.Context, datasetFunc func(context.Context, [][]float64), recoverFunc func(context.Context, error)) {
	defer func() {
		// 异常处理
		if exception := recover(); exception != nil {
			if recoverFunc != nil {
				recoverFunc(ctx, util.Error(exception))
			}
		}

		// 重启监听器
		gtimer.SetTimeout(ctx, detector.config.RestartInterval, func(context.Context) {
			detector.Listener(ctx, datasetFunc, recoverFunc)
		})
	}()

	// 端口列表
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		panic(err)
	}

	// 端口名称
	var portName string
	for _, port := range ports {
		// 匹配设备vid和pid
		if port.VID == detector.config.Vid && port.PID == detector.config.Pid {
			portName = port.Name
			break
		}
	}
	if portName == "" {
		panic("Serial port not found")
	}

	// 打开端口
	port, err := serial.Open(portName, &serial.Mode{
		BaudRate: detector.config.BaudRate,
	})
	if err != nil {
		panic(err)
	}

	// 数据集
	dataset := make([][]float64, 0, gconv.Int(math.Ceil(gconv.Float64(time.Second.Milliseconds())/gconv.Float64(detector.config.DataInterval))))

	for {
		// 缓冲区
		buffer := make([]byte, 64)

		// 读取数据
		size, err := port.Read(buffer)
		if err != nil {
			panic(err)
		}

		// 解析缓冲区
		data := detector.parseBuffer(buffer[:size])
		if len(data) > 0 {
			dataset = append(dataset, data)
		}

		// 处理数据集
		if len(dataset) == cap(dataset) {
			if datasetFunc != nil {
				datasetFunc(ctx, dataset)
			}
			dataset = dataset[:0]
		}
	}

	return
}

// StandbyData 待机数据
func (detector *Detector) StandbyData() []float64 {
	return detector.config.StandbyData
}

// parseBuffer 解析缓冲区
func (detector *Detector) parseBuffer(buffer []byte) (data []float64) {
	// 数据大小
	dataSize := detector.config.DataSize
	if dataSize != len(buffer) {
		return
	}

	// 数据帧头
	dataHead := detector.config.DataHead
	dataHeadSize := len(dataHead)
	for index, value := range buffer[:dataHeadSize] {
		if value != dataHead[index] {
			return
		}
	}

	// 数据帧尾
	dataTail := detector.config.DataTail
	dataTailSize := len(dataTail)
	for index, value := range buffer[dataSize-dataTailSize:] {
		if value != dataTail[index] {
			return
		}
	}

	// 提取数据
	data = make([]float64, 0, dataSize-dataHeadSize-dataTailSize)
	for _, value := range buffer[dataHeadSize : dataSize-dataTailSize] {
		data = append(data, gconv.Float64(value)/16)
	}

	return
}
