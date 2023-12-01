package main

import (
	"fmt"

	"github.com/4rchr4y/gouml/plantuml/ast"
	"github.com/4rchr4y/gouml/plantuml/types"
)

func main() {
	cs := ast.ClassSpec{
		Name:       "Hello",
		Visibility: types.Private,
	}

	ts := cs.Tokenify()

	fmt.Println(cs.Plantify(ts))
}
