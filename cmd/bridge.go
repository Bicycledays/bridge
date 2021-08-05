package main

/*
1200 бит/с
паритет нечётный
7 бит
1 стоповый
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"os"
	"strings"
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
	var measure []string

	for {
		n, err := s.Read(buf)

		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%q", buf[:n])

		if n > 0 {
			fmt.Print(buf[:n])
			measure = append(measure, fmt.Sprintf("%q", buf[0]))

			if buf[0] == 10 {
				//ch <- strings.Join(measure, "")
				sendMeasure(strings.Join(measure, ""))
				fmt.Println("")
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

func sendMeasure(measure string) {
	url, exists := os.LookupEnv("URL")

	if exists {

		body := map[string]string{"measure": measure}
		jsonData, err := json.Marshal(body)

		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(url, "application/json",
			bytes.NewBuffer(jsonData))

		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["success"])
	}
}
