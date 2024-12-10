package detector

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	patternVid             = "detector.vid"             // 配置供应商id
	patternPid             = "detector.pid"             // 配置产品id
	patternBaudRate        = "detector.baudRate"        // 配置波特率
	patternDataInterval    = "detector.dataInterval"    // 配置数据间隔，单位：毫秒
	patternDataSize        = "detector.dataSize"        // 配置数据大小
	patternDataHead        = "detector.dataHead"        // 配置数据帧头
	patternDataTail        = "detector.dataTail"        // 配置数据帧尾
	patternRestartInterval = "detector.restartInterval" // 配置重启间隔
	patternStandbyData     = "detector.standbyData"     // 配置待机数据
	defaultVid             = "1A86"                     // 默认供应商id
	defaultPid             = "7523"                     // 默认产品id
	defaultBaudRate        = 115200                     // 默认波特率
	defaultDataInterval    = 100                        // 默认数据间隔，单位：毫秒
	defaultDataSize        = 14                         // 默认数据大小
	defaultRestartInterval = "5s"                       // 默认重启间隔
)

type Config struct {
	Vid             string        // 供应商id
	Pid             string        // 产品id
	BaudRate        int           // 波特率
	DataInterval    int           // 数据间隔，单位：毫秒
	DataSize        int           // 数据大小
	DataHead        []byte        // 数据帧头
	DataTail        []byte        // 数据帧尾
	RestartInterval time.Duration // 重启间隔
	StandbyData     []float64     // 待机数据
}

// Init 初始化配置
func (config *Config) Init() {
	ctx := context.TODO()

	if config.Vid = g.Config().MustGet(ctx, patternVid).String(); config.Vid == "" {
		config.Vid = defaultVid
	}
	if config.Pid = g.Config().MustGet(ctx, patternPid).String(); config.Pid == "" {
		config.Pid = defaultPid
	}
	if config.BaudRate = g.Config().MustGet(ctx, patternBaudRate).Int(); config.BaudRate <= 0 {
		config.BaudRate = defaultBaudRate
	}
	if config.DataInterval = g.Config().MustGet(ctx, patternDataInterval).Int(); config.DataInterval <= 0 {
		config.DataInterval = defaultDataInterval
	}
	if config.DataSize = g.Config().MustGet(ctx, patternDataSize).Int(); config.DataSize <= 0 {
		config.DataSize = defaultDataSize
	}
	if config.DataHead = g.Config().MustGet(ctx, patternDataHead).Bytes(); len(config.DataHead) == 0 {
		config.DataHead = []byte{255, 129}
	}
	if config.DataTail = g.Config().MustGet(ctx, patternDataTail).Bytes(); len(config.DataTail) == 0 {
		config.DataTail = []byte{204, 90}
	}
	if config.RestartInterval = g.Config().MustGet(ctx, patternRestartInterval).Duration(); config.RestartInterval <= 0 {
		config.RestartInterval = gconv.Duration(defaultRestartInterval)
	}
	if config.StandbyData = g.Config().MustGet(ctx, patternStandbyData).Float64s(); len(config.StandbyData) == 0 {
		config.StandbyData = []float64{10.5, 7.8125, 7.8125, 7.1875, 7.875, 7.625, 7.875, 7.875, 7.875, 10.125}
	}
}
