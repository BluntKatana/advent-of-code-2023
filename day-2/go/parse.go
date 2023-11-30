package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func parse(fileName string) []string {
	// Open file
	file, err := os.Open(fileName)

	// Check for errors when opening file
	if err != nil {
		log.Fatal(err)
	}
	// Close file when done
	defer file.Close()

	// Read file
	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Split file into lines
	lines := strings.Split(string(input), "\n")

	return lines
}
