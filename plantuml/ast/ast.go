package ast

import (
	"github.com/4rchr4y/gouml/plantuml/token"
	"github.com/4rchr4y/gouml/plantuml/types"
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

type Spec interface {
	Connectable

	specElem() (token.Lexeme, token.Token)
	visibility() types.VisibilityKind
}

type ClassSpec struct {
	Name       string
	Visibility types.VisibilityKind
	// Type
}

func (d *ClassSpec) Tokenify() *token.TokenSet {
	tokenSet := token.NewTokenSet()

	tokenSet.AddToken(d.Visibility.Token())
	tokenSet.AddToken(token.CLASS)
	tokenSet.AddLexeme(d.Name)
	tokenSet.AddToken(types.Bracket.Opening())
	tokenSet.AddToken(token.NL)
	tokenSet.AddToken(types.Bracket.Closing())

	return tokenSet
}

func (d *ClassSpec) Plantify(tokSet *token.TokenSet) string {
	var result string

	for next, ok := tokSet.Next(); ok; next, ok = tokSet.Next() {
		result += next
	}

	return result
}

type ClassType struct {
	// GenericArgs
}
