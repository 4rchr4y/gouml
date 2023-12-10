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

func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type (
	// Token token is an identifier for various syntactic features supported by PlantUML.
	Token int

	// TokenSet is a result set of tokenization of any entity within a defined PlantUML AST.
	TokenSet struct {
		TokSeq        []Token  // sequence of tokens, e.g. `IDENTIFIER``
		LexSeq        []string // sequence of tokens lexical values, e.g. `MyFunc`
		TestSeq       int
		TestSeqStruct struct {
			TokSeq  []Token  // sequence of tokens, e.g. `IDENTIFIER``
			LexSeq  []string // sequence of tokens lexical values, e.g. `MyFunc`
			TestSeq int
		}
		TestSeqRef *Connectable
	}
	Vector[T int] []T
)

type testtype_a int
type testtype_b []int
type testtype_c *Connectable
type testtype_d chan int

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

func asd(a int, b string) {}
