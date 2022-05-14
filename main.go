package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	message := flag.String("m", "", "message input (in priority if other inputs chosen)")
	outputFilename := flag.String("o", "unknown.png", "output filename")
	inputFilename := flag.String("i", "", "input filename (.txt only)")
	flag.Parse()

	payload := ""

	switch {
	case *message != "":
		payload = *message
	case *inputFilename != "":
		if splitParts := strings.Split(*inputFilename, "."); len(splitParts) == 0 || splitParts[len(splitParts)-1] != "txt" {
			log.Fatal("wrong file extension")
		}

		bytes, err := os.ReadFile(*inputFilename)
		if err != nil {
			log.Fatalf("cannot read file %s: %v", *inputFilename, err)
		}
		payload = string(bytes)
	default:
		log.Fatal("no any input selected, check '-h' flag")
	}

	err := qrcode.WriteFile(payload, qrcode.Medium, 256, *outputFilename)
	if err != nil {
		log.Fatalf("cannot write file: %v", err)
	}
}
