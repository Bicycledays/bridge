package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	buf := make([]byte, 1024)
	for i := range buf {
		buf[i] = 8
	}

	n := 20
	log.Printf("%d", buf[n])
}
