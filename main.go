package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func readPayload(message, inputFilename string) (string, error) {
	if message != "" {
		return message, nil
	}

	if inputFilename == "" {
		return "", errors.New("no any input selected, check '-h' flag")
	}

	if splitParts := strings.Split(inputFilename, "."); len(splitParts) == 0 || splitParts[len(splitParts)-1] != "txt" {
		return "", errors.New("wrong file extension, check '-h' flag")
	}

	bytes, err := os.ReadFile(inputFilename)
	if err != nil {
		return "", fmt.Errorf("cannot read file %s: %w", inputFilename, err)
	}

	return string(bytes), nil
}

func main() {
	message := flag.String("m", "", "message input (in priority if other inputs chosen)")
	outputFilename := flag.String("o", "unknown.png", "output filename")
	inputFilename := flag.String("i", "", "input filename (.txt only)")
	flag.Parse()

	payload, err := readPayload(*message, *inputFilename)
	if err != nil {
		log.Fatalf("cannot read input: %v", err)
	}

	err = qrcode.WriteFile(payload, qrcode.Medium, 256, *outputFilename)
	if err != nil {
		log.Fatalf("cannot write file: %v", err)
	}
}
