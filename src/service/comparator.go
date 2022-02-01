package service

import (
	"github.com/tarm/serial"
	"log"
	"strconv"
	"time"
)

type Comparator struct {
	Config      *serial.Config `json:"config"`
	Params      *Params        `json:"params"`
	Display     []byte
	Subscribers int
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
		Config: nil,
		Params: nil,
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

func (c *Comparator) SendWhileListing(p *serial.Port) {
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	var err error
	for c.Subscribers > 0 {
		err = c.Send(p, Print)
		if err != nil {
			log.Println(err.Error())
		}
		<-ticker.C
	}
}

func (c *Comparator) Listen(p *serial.Port) {
	buf := make([]byte, 1)
	var measure []byte

	for c.Subscribers > 0 {
		_, err := p.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if buf[0] == '\n' {
			log.Println(string(measure))
			c.Display = measure
			measure = nil
		} else {
			measure = append(measure, buf[0])
		}
	}

	c.Display = nil
}

func (c *Comparator) isValidKey() bool {
	log.Println(c)
	if c.Params == nil {
		return false
	}
	s := c.Params.Number + "%" + strconv.Itoa(c.Params.Id) + "%" + c.Params.Term
	k := encrypt(s)
	log.Println(k)
	log.Println(c.Params.Key)
	return k == c.Params.Key
}
