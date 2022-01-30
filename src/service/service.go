package service

import (
	"errors"
	"github.com/tarm/serial"
)

type ComparatorService interface {
	OpenPort() *serial.Port
	Send(p *serial.Port, code Code)
	Listen(ch chan string, p *serial.Port)
	isValidKey() bool
}

type Scanner interface {
	RefreshPorts()
	GetPorts() []*Port
}

type Service struct {
	Comparators map[string]*Comparator
	Scanner     Scanner
}

func NewService() *Service {
	return &Service{
		Comparators: make(map[string]*Comparator),
		Scanner:     NewComputer(),
	}
}

func (s *Service) CheckComparator(c *Comparator) (portName string, err error) {
	if !c.isValidKey() {
		return "", errors.New("license key is not valid")
	}
	portName = c.Config.Name
	_, ok := s.Comparators[portName]
	if !ok {
		s.Comparators[c.Config.Name] = c
	}
	return portName, nil
}
