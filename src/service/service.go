package service

import (
	"encoding/json"
	"errors"
	"github.com/tarm/serial"
	"log"
)

type ComparatorService interface {
	OpenPort() *serial.Port
	Send(p *serial.Port, command Command)
	Listen(ch chan string, p *serial.Port)
	isValidKey() bool
}

type Scanner interface {
	RefreshPorts() error
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

func (s *Service) CheckComparator(js []byte) (hash string, err error) {
	var c Comparator
	err = json.Unmarshal(js, &c)
	if err != nil {
		return "", errors.New("comparator parsing error:" + err.Error())
	}
	if !c.isValidKey() {
		return "", errors.New("license key is not valid")
	}
	hash, err = c.hashConfig()
	if err != nil {
		return "", err
	}
	_, ok := s.Comparators[hash]
	if !ok {
		log.Println("new hash", hash)
		s.Comparators[hash] = &c
	} else {
		log.Println("existed hash", hash)
	}
	return hash, nil
}
