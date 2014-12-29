package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

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

func identifyType(decl ast.Decl) (typeName string, match bool) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	if genDecl.Doc == nil {
		return
	}

	found := false
	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "go:generate joiner") {
			found = true
			break
		}
	}
	if !found {
		return
	}

	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if typeSpec.Name != nil {
				typeName = typeSpec.Name.Name
				break
			}
		}
	}
	if typeName == "" {
		return
	}

	match = true
	return
}

func identifyStringer(decl ast.Decl) (typeName string, match bool) {
	funcDecl, ok := decl.(*ast.FuncDecl)
	if !ok {
		return
	}

	// Method name should match fmt.Stringer
	if funcDecl.Name == nil {
		return
	}
	if funcDecl.Name.Name != "String" {
		return
	}

	// Should have no arguments
	if funcDecl.Type == nil {
		return
	}
	if funcDecl.Type.Params == nil {
		return
	}
	if len(funcDecl.Type.Params.List) != 0 {
		return
	}

	// Return value should be a string
	if funcDecl.Type.Results == nil {
		return
	}
	if len(funcDecl.Type.Results.List) != 1 {
		return
	}
	result := funcDecl.Type.Results.List[0]
	if result.Type == nil {
		return
	}
	if ident, ok := result.Type.(*ast.Ident); !ok {
		return
	} else if ident.Name != "string" {
		return
	}

	// Receiver type
	if funcDecl.Recv == nil {
		return
	}
	if len(funcDecl.Recv.List) != 1 {
		return
	}
	recv := funcDecl.Recv.List[0]
	if recv.Type == nil {
		return
	}
	if ident, ok := recv.Type.(*ast.Ident); !ok {
		return
	} else {
		typeName = ident.Name
	}

	match = true
	return
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	log.Println(spew.Sdump(f.Decls))

	types := map[string]bool{}
	stringers := map[string]bool{}

	for _, decl := range f.Decls {
		typeName, ok := identifyType(decl)
		if ok {
			types[typeName] = true
			continue
		}

		typeName, ok = identifyStringer(decl)
		if ok {
			stringers[typeName] = true
			continue
		}
	}
	log.Printf("Types are %#v", types)
	log.Printf("Stringers are %#v", stringers)
}
