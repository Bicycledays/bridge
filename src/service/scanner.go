package service

import (
	"fmt"
	"github.com/hedhyw/Go-Serial-Detector/pkg/v1/serialdet"
	"log"
	"runtime"
	"strconv"
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

func (c *Computer) RefreshPorts() {
	scanner := c.getScanner()
	c.Ports = scanner()
}

func (c *Computer) getScanner() func() []*Port {
	var scanner func() []*Port

	switch c.System {
	case "windows":
		scanner = c.scanWindows
	case "linux":
		scanner = c.scanLinux
	case "ios":
		scanner = c.scanMac
	default:
		panic(fmt.Sprintf("Операционая система %s", c.System))
	}

	return scanner
}

func (c *Computer) scanWindows() []*Port {
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
	return make([]*Port, 0)
}

func (c *Computer) scanLinux() []*Port {
	var port Port
	var ports []*Port

	list, err := serialdet.List()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range list {
		port = Port{Name: p.Path(), IsBusy: false}
		log.Println(port)
		ports = append(ports, &port)
	}

	index := len(ports)
	ports = append(ports, &Port{
		Name:   "/dev/ttyUSB" + strconv.Itoa(index),
		IsBusy: false,
	})

	return ports
}

func (c *Computer) scanMac() []*Port {
	// todo
	return make([]*Port, 0)
}
