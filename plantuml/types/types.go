package types

import "github.com/4rchr4y/gouml/plantuml/token"

type VisibilityKind int

const (
	Public         = token.PLUS
	Private        = token.MINUS
	Protected      = token.POUND
	PackagePrivate = token.TILDE
)

type DelimiterKind int

const (
	Parenthesis = iota // ( ... )
	Brace              // [ ... ]
	Bracket            // { ... }
	Invisible
)
