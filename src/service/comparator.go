package service

import (
	"encoding/json"
	"github.com/tarm/serial"
	"log"
)

type Comparator struct {
	Config  *serial.Config `json:"config"`
	License *License       `json:"license"`
}

type Route string
type Code byte

const (
	RouteTare     Route = "/tare"
	RoutePrint    Route = "/print"
	RouteCover    Route = "/manage-cover"
	RoutePlatform Route = "/manage-platform"
)

const (
	CodeEsc      Code = 27  // ESC
	CodeTare     Code = 'T' // тарирование или установка на ноль
	CodePrint    Code = 'P' // печать
	CodeCover    Code = 'X'
	CodePlatform Code = 'Y'
)

func (c *Comparator) OpenPort() *serial.Port {
	p, err := serial.OpenPort(c.Config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return p
}

/*
	Передача команды на компаратор
*/
func (c *Comparator) Send(p *serial.Port, code Code) {
	buf := make([]byte, 2)
	buf[0] = byte(CodeEsc)
	buf[1] = byte(code)
	_, err := p.Write(buf)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (c *Comparator) Listen(p *serial.Port) {
	buf := make([]byte, 1)
	var measure []byte

	for {
		_, err := p.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if buf[0] == '\n' {
			log.Println("measure " + string(measure))
			measure = nil
		} else {
			measure = append(measure, buf[0])
		}
	}
}

func (c *Comparator) isValidKey() bool {
	l := make(map[string]string, 3)
	l["model"] = c.License.Model
	l["factoryNumber"] = c.License.Number
	l["licenseTerm"] = c.License.Term
	js, err := json.Marshal(l)
	if err != nil {
		return false
	}
	k := encrypt(string(js))
	return k != c.License.Key
}
