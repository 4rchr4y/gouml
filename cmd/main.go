package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/4rchr4y/gouml/plantuml/uml"
)

func main() {
	args := os.Args
	pathToFile := args[0]

	if pathToFile == "" {
		pathToFile = "./plantuml/ast/ast.go"
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "plantuml/ast/ast.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(node)
	uml.AstToPlantuml(node)

	os.Exit(1)
}
