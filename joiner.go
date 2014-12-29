// Copyright 2014 Brett Slatkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	defer output.Close()

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
