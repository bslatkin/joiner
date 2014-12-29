package main

import (
	"go/parser"
	"go/token"
	"log"

	"github.com/davecgh/go-spew/spew"
)

var src = `
package main

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
	log.Println(spew.Sdump(f.Decls))

	joiners := map[string]bool{}
	stringers := map[string]bool{}

	for _, decl := range f.Decls {
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
	log.Printf("Joiners are %#v", joiners)
	log.Printf("Stringers are %#v", stringers)
}
