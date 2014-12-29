package main

import (
	"log"
	"os"
)

func processFile(inputPath string) {
	log.Printf("Processing file %s", inputPath)

	packageName, types := loadFile(inputPath)

	log.Printf("Found joiner types to generate: %#v", types)

	outputPath, err := getRenderedPath(inputPath)
	if err != nil {
		log.Fatalf("Could not get output path: %s", err)
	}

	output, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("Could not open output file: %s", err)
	}

	if err := render(output, packageName, types); err != nil {
		log.Fatalf("Could not generate go code: %s", err)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("joiner: ")

	for _, path := range os.Args[1:] {
		processFile(path)
	}
}
