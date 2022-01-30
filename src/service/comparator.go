package service

import (
	"encoding/json"
	"github.com/tarm/serial"
	"log"
)

type Comparator struct {
	Config  *serial.Config  `json:"config"`
	License *LicenseService `json:"license"`
}

type Code byte

const (
	Esc      Code = 27  // ESC
	Tare     Code = 'T' // тарирование или установка на ноль
	Print    Code = 'P' // печать
	Cover    Code = 'X'
	Platform Code = 'Y'
)

func NewComparator() *Comparator {
	return &Comparator{
		Config:  nil,
		License: nil,
	}
}

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
func (c *Comparator) Send(p *serial.Port, code Code) error {
	buf := make([]byte, 2)
	buf[0] = byte(Esc)
	buf[1] = byte(code)
	_, err := p.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comparator) Listen(ch chan []byte, p *serial.Port) {
	buf := make([]byte, 1)
	var measure []byte

	for {
		_, err := p.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if buf[0] == '\n' {
			log.Println("write inside channel")
			select {
			case ch <- measure:
				measure = nil
			default:
				return
			}
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
	log.Println(k != c.License.Key)
	return k == c.License.Key
}
