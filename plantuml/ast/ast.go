package ast

import "github.com/4rchr4y/gouml/plantuml/types"

type (
	Token int

	TokenSet struct {
		TokSeq        []Token
		LexSeq        []string
		TestSeq       int
		TestSeqStruct *Connectable
	}
) //TODO здесь вот этой конструкции не должно быть, она просто для тестов

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
