package main

import (
	"fmt"
	"log"

	"go.bug.st/serial/enumerator"
)

type Device struct{}

func (*Device) listener() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return
	}
	for _, port := range ports {
		fmt.Printf("Port: %s\n", port.Name)
		if port.Product != "" {
			fmt.Printf("   Product Name: %s\n", port.Product)
		}
		if port.IsUSB {
			fmt.Printf("   USB ID      : %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial  : %s\n", port.SerialNumber)
		}
	}
}
