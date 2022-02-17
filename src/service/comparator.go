package service

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

type Comparator struct {
	Config      *serial.Config `json:"config"`
	Params      *Params        `json:"params"`
	Display     []byte         `json:"-"`
	Subscribers int            `json:"-"`
}

func NewComparator() *Comparator {
	return &Comparator{
		Config: nil,
		Params: nil,
	}
}

func (c *Comparator) OpenPort() (*serial.Port, error) {
	p, err := serial.OpenPort(c.Config)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Send Передача команды на компаратор
func (c *Comparator) Send(p *serial.Port, command Command) error {
	_, err := p.Write(command.message())
	if err != nil {
		return err
	}
	return nil
}

func (c *Comparator) SendWhileListing(p *serial.Port) {
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	command := Command{
		Format: 1,
		Symbol: Print,
	}

	for c.Subscribers > 0 {
		err := c.Send(p, command)
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
			log.Println(err)
			break
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
	log.Println(*c)
	if c.Params == nil {
		return false
	}
	s := c.Params.Number + "%" + c.Params.Model + "%" + c.Params.Term
	k := encrypt(s)
	log.Println(k)
	log.Println(c.Params.Key)
	log.Println(k == c.Params.Key)
	return k == c.Params.Key
}
