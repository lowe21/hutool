package main

import (
	"log"
	"time"

	"go.bug.st/serial"

	"hutool/internal/pkg/websocket"
)

type Device struct{}

func (device *Device) listener() {
	port, err := serial.Open("", &serial.Mode{
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
			log.Printf("%v", data[:n])
			log.Printf("%s", data[:n])
		}
	}

	<-time.After(3 * time.Second)

	device.listener()
}
