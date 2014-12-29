package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
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
	fmt.Printf("%#v\n", f.Decls)

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Doc != nil {
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "go:generate joiner") {
						fmt.Printf("Found joiner on %#v\n", genDecl)
					}
				}
			}
		}
	}
}
