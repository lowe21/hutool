package main

import (
	"fmt"
	"time"

	"go.bug.st/serial"

	"github.com/gogf/gf/v2/text/gstr"

	"hutool/internal/pkg/websocket"
)

type Device struct{}

func (device *Device) listener() {
	// VID:PID 1A86:7523
	port, err := serial.Open("/dev/cu.usbserial-1460", &serial.Mode{
		BaudRate: 115200,
	})
	if err != nil {
		websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
	} else {
		count := 0
		data := make([]map[string]any, 0, 10)

		for {
			buffer := make([]byte, 128)
			n, err := port.Read(buffer)
			if err != nil {
				websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
				break
			}

			buffer = device.parse(buffer[:n])
			if len(buffer) == 0 {
				continue
			}

			for index, value := range buffer {
				data = append(data, map[string]any{
					"line": gstr.JoinAny([]any{"line", index}, ""),
					"x":    count,
					"y":    value,
				})
			}

			if count++; count > 10 {
				websocket.Notice(websocket.Message("data", data))
				count = 0
				data = data[:0]
			}
		}
	}

	<-time.After(5 * time.Second)

	device.listener()
}

func (*Device) parse(buffer []byte) []byte {
	length := len(buffer)
	if length != 14 {
		return nil
	}

	header := make([]string, 0, 2)
	for _, value := range buffer[:2] {
		header = append(header, fmt.Sprintf("%X", value))
	}
	if gstr.Join(header, "") != "FF81" {
		return nil
	}

	footer := make([]string, 0, 2)
	for _, value := range buffer[length-2:] {
		footer = append(footer, fmt.Sprintf("%X", value))
	}
	if gstr.Join(footer, "") != "CC5A" {
		return nil
	}

	return buffer[2 : length-2]
}
