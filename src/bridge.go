package main

/*
1200 бит/с
паритет нечётный
7 бит
1 стоповый
*/

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tarm/serial"
	"log"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	c := &serial.Config{
		Name:     "COM4",
		Baud:     1200,
		Parity:   serial.ParityOdd,
		Size:     7,
		StopBits: serial.Stop1,
	}

	s, err := serial.OpenPort(c)
	ch := make(chan bool)

	if err != nil {
		log.Fatal(err)
	}

	go listen(s, ch)
	go write(s)

	<-ch
}

func listen(s *serial.Port, ch chan bool) {

	buf := make([]byte, 1)
	var measure []byte

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			if buf[0] == 10 {
				fmt.Println(string(measure))
				measure = nil
			} else {
				measure = append(measure, buf[0])
			}
		}
	}

	ch <- true
}

func write(s *serial.Port) {
	data := make([]byte, 20)
	data[0] = 0x1B
	data[1] = 0x50

	for {
		n, err := s.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(n)
		time.Sleep(10 * time.Second)
	}
}
