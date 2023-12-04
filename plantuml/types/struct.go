package types

import (
	"go/ast"
	"go/token"
	"unicode"
)

type StructModel struct {
	Name       ident
	Type       *structModelTypes
	Visibility VisibilityKind
	style      *structStyle
	extends    string
	Field      *[]structFieldAndMethod
	Methods    *[]structFieldAndMethod
}

type structStyleLineType string

const (
	bold   structStyleLineType = "bold"
	dashed structStyleLineType = "dashed"
	dotted structStyleLineType = "dotted"
)

type structStyle struct {
	text       string // color name or hex(without #)
	background string
	header     string
	line       string
	lineType   structStyleLineType
}

type ident struct {
	Name  string
	Title string
}

type structModelTypes string

const (
	AbstractClass structModelTypes = "abstract class"
	Annotation    structModelTypes = "annotation"
	Entity        structModelTypes = "entity"
	Enum          structModelTypes = "enum"
	Interface     structModelTypes = "interface"
	Protocol      structModelTypes = "protocol"
	Struct        structModelTypes = "struct"
	Exception     structModelTypes = "exception"
	Metaclass     structModelTypes = "metaclass"
	Stereotype    structModelTypes = "stereotype"
)

type structFieldAndMethod struct {
	Name       string
	Visibility VisibilityKind
	Type       structFieldAndMethodType
}

type structFieldAndMethodType string

const (
	Static   structFieldAndMethodType = "static"
	Abstract structFieldAndMethodType = "abstract"
	None     structFieldAndMethodType = ""
)

func firstLetterIsLowercase(s string) bool {
	if len(s) == 0 {
		return false
	}

	firstChar := rune(s[0])

	return unicode.IsLower(firstChar)
}

func getTypeSpecFieldStructs(Expr ast.Expr) *[]structFieldAndMethod {
	var Structs []structFieldAndMethod

	switch Type := Expr.(type) {
	case *ast.Ident:
		Structs = append(Structs, structFieldAndMethod{
			Name: Type.Name,
		})

	case *ast.StructType:
		for _, field := range Type.Fields.List {
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					Structs = append(Structs, structFieldAndMethod{
						Name: name.Name,

						Type: structFieldAndMethodType(None),
					})
				}
			}
		}
	case *ast.StarExpr:
		if ident, ok := Type.X.(*ast.Ident); ok {
			Structs = append(Structs, structFieldAndMethod{
				Name: ident.Name,
			})
		}
	}

	return &Structs
}

func getTypeSpecStruct(TypeSpec *ast.TypeSpec) StructModel {
	Type := Exception
	Struct := StructModel{
		Name: ident{
			Name: TypeSpec.Name.Name,
			// Title: TypeSpec.Name.Name,
		},
		Type: &Type, //TODO а как, блядь, сюда просто вставить Exception???
	}

	if firstLetterIsLowercase(TypeSpec.Name.Name) {
		Struct.Visibility = VisibilityKind(Private)
		Struct.style = &structStyle{
			lineType: structStyleLineType(dotted),
		}
	} else {
		Struct.Visibility = VisibilityKind(Public)
	}

	Struct.Field = getTypeSpecFieldStructs(TypeSpec.Type)

	return Struct
}

func GetStructFromDecl(decl *ast.GenDecl) []StructModel {
	var Structs []StructModel
	// if decl.Tok == token.IMPORT {

	// }

	if decl.Tok == token.TYPE {
		for _, Spec := range decl.Specs {
			Structs = append(Structs, getTypeSpecStruct(Spec.(*ast.TypeSpec)))
		}
	}

	return Structs
}
