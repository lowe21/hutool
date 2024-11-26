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
		for {
			data := make([]byte, 128)
			n, err := port.Read(data)
			if err != nil {
				websocket.Notice(websocket.Message("error", "DEVICE_ERROR", err.Error()))
				break
			}
			fmt.Printf("%X\n", data[:n])
			fmt.Printf("%v\n", data[:n])
		}
	}

	<-time.After(3 * time.Second)

	device.listener()
}
