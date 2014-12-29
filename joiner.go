package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var src = `
package mycrazypackage

import (
    "fmt"
    "strings"

    _ "github.com/bslatkin/joiner"
)

// go:generate joiner
type Person struct {
    FirstName string
    LastName  string
    HairColor string
}

// go:generate joiner
type MyType int

func (s Person) String() string {
    return fmt.Sprintf("%#v", s)
}

func main() {
    people := []Person{
        Person{"Sideshow", "Bob", "red"},
        Person{"Homer", "Simpson", "n/a"},
        Person{"Lisa", "Simpson", "blonde"},
        Person{"Marge", "Simpson", "blue"},
        Person{"Mr", "Burns", "gray"},
    }
    fmt.Printf("My favorite Simpsons Characters:%s\n", JoinPerson(people).With("\n"))
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
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

	if err := render(os.Stdout, packageName, types); err != nil {
		log.Fatalf("Could not generate go code: %s", err)
	}
}
