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

	for {
		size := 128
		buffer := make([]byte, size)

		size, err = port.Read(buffer)
		if err != nil {
			break
		}

		data := device.parse(buffer[:size])
		if len(data) > 0 {
			websocket.Notice(websocket.Message("data", device.base(data)))
		}
	}
}

func (*Device) parse(buffer []byte) (data []float64) {
	header := make([]string, 0, 2)
	for _, value := range buffer[:2] {
		header = append(header, fmt.Sprintf("%X", value))
	}
	if gstr.Join(header, "") != "FF81" {
		return nil
	}

	footer := make([]string, 0, 2)
	for _, value := range buffer[len(buffer)-2:] {
		footer = append(footer, fmt.Sprintf("%X", value))
	}
	if gstr.Join(footer, "") != "CC5A" {
		return nil
	}

	data = make([]float64, 0, len(buffer)-4)

	for _, value := range buffer[2 : len(buffer)-2] {
		data = append(data, gconv.Float64(value)/16)
	}

	return
}

func (*Device) base(data []float64) []map[string]any {
	line := map[int]float64{
		0: 10.5,
		1: 7.8125,
		2: 7.8125,
		3: 7.1875,
		4: 7.875,
		5: 7.625,
		6: 7.875,
		7: 7.875,
		8: 7.875,
		9: 10.125,
	}

	lines := make([]map[string]any, 0, 11*len(data))
	for index, value := range data {
		for x := range 11 {
			y := 0.0
			if x == 5 {
				y = value - line[index]
				if y <= 0.0625 && y >= -0.0625 {
					y = 0.0
				}
			}
			lines = append(lines, map[string]any{
				"line": gstr.JoinAny([]any{"line", index}, ""),
				"x":    x,
				"y":    y,
			})
		}
	}

	return lines
}
