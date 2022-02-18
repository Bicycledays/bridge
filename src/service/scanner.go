package service

import (
	"errors"
	"fmt"
	"go.bug.st/serial.v1"
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
	ports, err := c.getScanner()
	if err != nil {
		return err
	}
	c.Ports = ports
	return nil
}

func (c *Computer) getScanner() ([]*Port, error) {
	var scanner func() ([]*Port, error)

	switch c.System {
	case "windows":
		scanner = c.scanWindows
	case "linux":
		scanner = c.scanLinux
	case "ios":
		scanner = c.scanMac
	default:
		return nil, errors.New(fmt.Sprintf("unidentified operating system %s", c.System))
	}

	return scanner()
}

func (c *Computer) scanWindows() ([]*Port, error) {
	// todo
	//k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\\DEVICEMAP\\SERIALCOMM`, registry.QUERY_VALUE)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer k.Close()
	//
	//ki, err := k.Stat()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Subkey %d ValueCount %d\n", ki.SubKeyCount, ki.ValueCount)
	//
	//s, err := k.ReadValueNames(int(ki.ValueCount))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//kvalue := make([]string, ki.ValueCount)
	//
	//for i, test := range s {
	//	q, _, err := k.GetStringValue(test)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	kvalue[i] = q
	//}
	//
	//fmt.Printf("%s \n", kvalue)
	return nil, nil
}

func (c *Computer) scanLinux() ([]*Port, error) {
	result, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	ports := filterPorts(result, "ttyUSB")

	return ports, nil
}

func (c *Computer) scanMac() ([]*Port, error) {
	result, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	ports := filterPorts(result, "tty.usbserial")

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
