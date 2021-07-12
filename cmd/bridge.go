package main

/*
1200 бит/с
паритет нечётный
7 бит
1 стоповый
*/

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func main() {

	c := &serial.Config{
		Name:     "COM4",
		Baud:     1200,
		Parity:   serial.ParityOdd,
		Size:     7,
		StopBits: serial.Stop1,
	}

	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	listen(s)
}

func listen(s *serial.Port) {
	buf := make([]byte, 1024)

	for {
		n, err := s.Read(buf)

		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%q", buf[:n])

		if n > 0 {
			fmt.Print(buf[:n])

			if buf[0] == 67 {
				fmt.Println("")
			}
		}
	}
}
