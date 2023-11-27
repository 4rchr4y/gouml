package types

import "github.com/4rchr4y/gouml/plantuml/token"

type VisibilityKind int

const (
	Public         = token.PUBLIC
	Private        = token.PRIVATE
	Protected      = token.PROTECTED
	PackagePrivate = token.PACKAGE_PRIVATE
)
