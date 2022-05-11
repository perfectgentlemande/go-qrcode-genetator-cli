package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

const (
	inputTypeJSON    = "json"
	inputTypeMessage = "message"
	inputTypeTXT     = "txt"
)

func main() {
	message := flag.String("m", "no message", "message input")
	filename := flag.String("o", "unknown.png", "output filename")
	inputType := flag.String("i", inputTypeMessage, "input type")

	flag.Parse()

	fmt.Println(*message)
	fmt.Println(*inputType)

	err := qrcode.WriteFile(*message, qrcode.Medium, 256, *filename)
	if err != nil {
		log.Fatalf("cannot write file: %v", err)
	}
}
