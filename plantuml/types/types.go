package types

import "github.com/4rchr4y/gouml/plantuml/token"

type VisibilityKind token.Token

const (
	Public         VisibilityKind = VisibilityKind(token.PLUS)
	Private        VisibilityKind = VisibilityKind(token.MINUS)
	Protected      VisibilityKind = VisibilityKind(token.POUND)
	PackagePrivate VisibilityKind = VisibilityKind(token.TILDE)
)

func (v VisibilityKind) Token() token.Token {
	return token.Token(v)
}

type BoundKind [2]token.Token

var (
	Parenthesis BoundKind = [2]token.Token{token.LPAREN, token.RPAREN} // ( ... )
	Brace       BoundKind = [2]token.Token{token.LBRACK, token.RBRACK} // [ ... ]
	Bracket     BoundKind = [2]token.Token{token.LBRACE, token.RBRACE} // { ... }
	Invisible   BoundKind = [2]token.Token{token.EMPTY, token.EMPTY}   //
)

func (b BoundKind) Opening() token.Token { return b[0] }
func (b BoundKind) Closing() token.Token { return b[1] }

func (b BoundKind) Wrap(tokSet []token.Token) []token.Token {
	result := make([]token.Token, 0, len(tokSet)+2)
	result = append(result, b.Opening())
	result = append(result, tokSet...)
	result = append(result, b.Closing())
	return result
}
