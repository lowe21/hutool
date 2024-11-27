package main

import (
	"fmt"
	"time"

	"go.bug.st/serial"

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
			if n != 14 {
				continue
			}

			var (
				header string
				footer string
			)
			for _, value := range data[:2] {
				header += fmt.Sprintf("%X", value)
			}
			for _, value := range data[n-2:] {
				footer += fmt.Sprintf("%X", value)
			}
			if header != "FF81" || footer != "CC5A" {
				continue
			}

			for index, value := range buffer[2 : n-2] {
				data = append(data, map[string]any{
					"line": index,
					"x":    count,
					"y":    value,
				})
			}

			count++

			if count >= 10 {
				websocket.Notice(websocket.Message("error", data))
				count = 0
				data = data[:0]
			}
		}
	}

	<-time.After(3 * time.Second)

	device.listener()
}
