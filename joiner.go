package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func processFile(inputPath string) {
	log.Printf("Processing file %s", inputPath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	packageName := identifyPackage(f)
	if packageName == "" {
		log.Fatal("Could not determine package name")
	}

	joiners := map[string]bool{}
	stringers := map[string]bool{}
	for _, decl := range f.Decls {
		log.Printf("Considering %s", spew.Sdump(decl))

		typeName, ok := identifyJoinerType(decl)
		if ok {
			joiners[typeName] = true
			continue
		}

		typeName, ok = identifyStringer(decl)
		if ok {
			stringers[typeName] = true
			continue
		}
	}

	types := []GeneratedType{}
	for typeName, _ := range joiners {
		_, isStringer := stringers[typeName]
		joiner := GeneratedType{typeName, isStringer}
		types = append(types, joiner)
	}

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

	processFile("./example/main.go")
}
