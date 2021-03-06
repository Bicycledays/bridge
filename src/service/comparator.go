package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

	err := p.Close()
	if err != nil {
		log.Println("notice:", err.Error())
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

	today := time.Now()
	term, err := time.Parse("2006-01-02", c.Params.Term)

	return err == nil && term.After(today) && k == c.Params.Key
}

func (c *Comparator) hashConfig() (string, error) {

	config, err := json.Marshal(c.Config)
	if err != nil {
		log.Println("hashing config error", err.Error())
		return "", err
	}

	hmd5 := md5.Sum(config)
	return hex.EncodeToString(hmd5[:]), nil
}
