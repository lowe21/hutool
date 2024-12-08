package main

import (
	"context"
	"fmt"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"

	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"hutool/internal/pkg/websocket"
	"hutool/internal/util"
)

var (
	standbyData = []float64{10.5, 7.8125, 7.8125, 7.1875, 7.875, 7.625, 7.875, 7.875, 7.875, 10.125}
)

type Device struct {
	VID      string
	PID      string
	BaudRate int
}

func (device *Device) listener() {
	var err error

	defer func() {
		if exception := recover(); exception != nil {
			err = util.Error(exception)
		}

		if err != nil {
			websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
		}

		gtimer.SetTimeout(context.TODO(), 5*time.Second, func(context.Context) {
			device.listener()
		})
	}()

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		err = util.Error(err)
		return
	}

	var portName string

	for _, port := range ports {
		if port.VID == device.VID && port.PID == device.PID {
			portName = port.Name
			break
		}
	}

	if portName == "" {
		err = util.Error("Serial port not found")
		return
	}

	port, err := serial.Open(portName, &serial.Mode{
		BaudRate: device.BaudRate,
	})
	if err != nil {
		err = util.Error(err)
		return
	}

	var (
		count    = 10
		dataList = make([][]float64, 0, count)
	)

	for {
		size := 64
		buffer := make([]byte, size)

		size, err = port.Read(buffer)
		if err != nil {
			break
		}

		data := device.parse(buffer[:size])
		if len(data) > 0 {
			count--
			dataList = append(dataList, data)
		}

		if count <= 0 {
			device.send(dataList)
			count = 10
			dataList = dataList[:0]
		}
	}
}

func (*Device) parse(buffer []byte) (data []float64) {
	if len(buffer) != 14 {
		return
	}

	header := make([]string, 0, 2)

	for _, value := range buffer[:2] {
		header = append(header, fmt.Sprintf("%X", value))
	}

	if gstr.Join(header, "") != "FF81" {
		return
	}

	footer := make([]string, 0, 2)

	for _, value := range buffer[len(buffer)-2:] {
		footer = append(footer, fmt.Sprintf("%X", value))
	}

	if gstr.Join(footer, "") != "CC5A" {
		return
	}

	data = make([]float64, 0, len(buffer)-4)

	for _, value := range buffer[2 : len(buffer)-2] {
		data = append(data, gconv.Float64(value)/16)
	}

	return
}

func (device *Device) send(dataList [][]float64) {
	data := make([]float64, 10)

	for _, item := range dataList {
		for index := range standbyData {
			data[index] += item[index]
		}
	}

	for index := range data {
		data[index] /= gconv.Float64(len(dataList))
	}

	newData := make([]map[string]any, 0, 30)
	x := 0

	for index := range standbyData {
		newData = append(newData, map[string]any{
			"line": gstr.JoinAny([]any{"line", index}, ""),
			"x":    x,
			"y":    0,
		})
		x++
	}

	switch {
	case device.isStandby(data):
		for index := range data {
			newData = append(newData, map[string]any{
				"line": gstr.JoinAny([]any{"line", index}, ""),
				"x":    x,
				"y":    0,
			})
			x++
		}
	default:
		for index, value := range data {
			newData = append(newData, map[string]any{
				"line": gstr.JoinAny([]any{"line", index}, ""),
				"x":    x,
				"y":    value,
			})
			x++
		}
	}

	for index := range standbyData {
		newData = append(newData, map[string]any{
			"line": gstr.JoinAny([]any{"line", index}, ""),
			"x":    x,
			"y":    0,
		})
		x++
	}

	websocket.Notice(websocket.Message("data", newData))
}

func (*Device) isStandby(data []float64) (result bool) {
	result = true

	for index, value := range standbyData {
		if !result {
			break
		}

		if data[index] != value {
			diff := data[index] - value
			if diff > 0.1 || diff < -0.1 {
				result = false
			}
		}
	}

	return
}
