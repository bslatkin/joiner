package main

import (
	"go/ast"
	"strings"
)

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
