package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/4rchr4y/gouml/plantuml/uml"
)

func main() {
	// args := os.Args

	pathToFile := "plantuml/test/test.go"

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, pathToFile, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	uml.AstToPlantuml(node)

	os.Exit(0)
}
