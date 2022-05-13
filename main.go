package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	message := flag.String("m", "no message", "message input")
	outputFilename := flag.String("o", "unknown.png", "output filename")
	inputFilename := flag.String("i", "input.txt", "input filename")

	flag.Parse()

	fmt.Println(*message)
	fmt.Println(*inputFilename)
	fmt.Println(*outputFilename)

	err := qrcode.WriteFile(*message, qrcode.Medium, 256, *inputFilename)
	if err != nil {
		log.Fatalf("cannot write file: %v", err)
	}
}
