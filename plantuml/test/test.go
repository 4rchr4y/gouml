package test

import (
	"github.com/4rchr4y/gouml/plantuml/types"
)

type (
	// Token token is an identifier for various syntactic features supported by PlantUML.
	Token int

	// TokenSet is a result set of tokenization of any entity within a defined PlantUML AST.
	TokenSet struct {
		TokSeq        []Token  // sequence of tokens, e.g. `IDENTIFIER``
		LexSeq        []string // sequence of tokens lexical values, e.g. `MyFunc`
		TestSeq       int
		TestSeqStruct *Connectable
	}
)

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

func asd() {}
