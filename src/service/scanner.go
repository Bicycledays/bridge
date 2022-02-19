package service

import (
	"errors"
	"fmt"
	"go.bug.st/serial.v1"
	"log"
	"runtime"
	"strings"
)

type Port struct {
	Name   string `json:"name"`
	IsBusy bool   `json:"isBusy"`
}

type Computer struct {
	Ports  []*Port `json:"ports"`
	System string  `json:"system"`
}

func NewComputer() *Computer {
	return &Computer{System: runtime.GOOS}
}

func (c *Computer) GetPorts() []*Port {
	return c.Ports
}

func (c *Computer) RefreshPorts() error {
	ports, err := c.ScanPorts()
	if err != nil {
		return err
	}
	c.Ports = ports
	return nil
}

func (c *Computer) ScanPorts() ([]*Port, error) {
	var sign string

	switch c.System {
	case "windows":
		sign = "COM"
	case "linux":
		sign = "ttyUSB"
	case "ios":
	case "darwin":
		sign = "tty.usbserial"
	default:
		return nil, errors.New(fmt.Sprintf("unidentified operating system %s", c.System))
	}

	return c.scan(sign)
}

func (c *Computer) scan(sign string) ([]*Port, error) {
	result, err := serial.GetPortsList()
	if err != nil {
		log.Println("scan")
		return nil, err
	}
	ports := filterPorts(result, sign)

	return ports, nil
}

func filterPorts(portsList []string, sign string) (ports []*Port) {
	for _, p := range portsList {
		if strings.Contains(p, sign) {
			ports = append(ports, &Port{
				Name:   p,
				IsBusy: true,
			})
		}
	}

	return ports
}
