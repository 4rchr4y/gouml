package ast

import (
	"math/cmplx"

	"github.com/4rchr4y/gouml/plantuml/types"
)

var a int
var (
	b      string
	c      int
	d                 = 4
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type (
	// Token token is an identifier for various syntactic features supported by PlantUML.
	Token int

	// TokenSet is a result set of tokenization of any entity within a defined PlantUML AST.
	TokenSet struct {
		TokSeq        []Token  // sequence of tokens, e.g. `IDENTIFIER``
		LexSeq        []string // sequence of tokens lexical values, e.g. `MyFunc`
		TestSeq       int
		TestSeqStruct struct{}
		TestSeqRef    Connectable
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
