package astParser

import "github.com/4rchr4y/gouml/plantuml/types"

type Element interface {
	// Tokenify
	// Plantify
}

type Connectable interface {
	Element
	connElem()
	// Connect(Connectable) Relationship
}

type Declaration interface {
	Connectable
	declElem()
}

type Ident struct {
	Name       string
	Visibility types.VisibilityKind
}

type ClassDecl struct{}
