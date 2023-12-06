package astParser

import (
	"go/ast"
	"unicode"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func firstLetterIsLowercase(s string) bool {
	if len(s) == 0 {
		return false
	}

	firstChar := rune(s[0])

	return unicode.IsLower(firstChar)
}

func getTypeSpecFieldStructs(Expr ast.Expr) *[]types.StructFieldAndMethod {
	var Structs []types.StructFieldAndMethod

	switch Type := Expr.(type) {
	case *ast.Ident:
		Structs = append(Structs, types.StructFieldAndMethod{
			Name: Type.Name,
		})

	case *ast.StructType:
		for _, field := range Type.Fields.List {
			for _, name := range field.Names {
				Structs = append(Structs, types.StructFieldAndMethod{
					Name: name.Name,
				})
			}
		}
	case *ast.StarExpr:
		if ident, ok := Type.X.(*ast.Ident); ok {
			Structs = append(Structs, types.StructFieldAndMethod{
				Name: ident.Name,
			})
		}
	case *ast.InterfaceType:
		for _, field := range Type.Methods.List {
			for _, name := range field.Names {
				Structs = append(Structs, types.StructFieldAndMethod{
					Name: name.Name,
				})
			}
		}
	}

	return &Structs
}

func GetTypeSpecStruct(TypeSpec ast.TypeSpec) types.StructModel {
	Type := types.Exception
	Struct := types.StructModel{
		Name: types.Ident{
			Name: TypeSpec.Name.Name,
			// Title: TypeSpec.Name.Name,
		},
		Type: &Type, //TODO а как, блядь, сюда просто вставить Exception???
	}

	if firstLetterIsLowercase(TypeSpec.Name.Name) {
		Struct.Visibility = types.VisibilityKind(types.Private)
		Struct.Style = &types.StructStyle{
			LineType: types.StructStyleLineType(types.Dotted),
		}
	} else {
		Struct.Visibility = types.VisibilityKind(types.Public)
	}

	Struct.Field = getTypeSpecFieldStructs(TypeSpec.Type)

	return Struct
}
