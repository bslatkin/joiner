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
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func loadFile(inputPath string) (string, []GeneratedType) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Could not parse file: %s", err)
	}

	packageName := identifyPackage(f)
	if packageName == "" {
		log.Fatalf("Could not determine package name of %s", inputPath)
	}

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

	types := []GeneratedType{}
	for typeName, _ := range joiners {
		_, isStringer := stringers[typeName]
		joiner := GeneratedType{typeName, isStringer}
		types = append(types, joiner)
	}

	return packageName, types
}

func identifyPackage(f *ast.File) string {
	if f.Name == nil {
		return ""
	}
	return f.Name.Name
}

func identifyJoinerType(decl ast.Decl) (typeName string, match bool) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	if genDecl.Doc == nil {
		return
	}

	found := false
	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@joiner") {
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
