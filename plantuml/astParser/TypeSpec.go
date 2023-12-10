package astParser

import (
	"go/ast"

	"github.com/4rchr4y/gouml/plantuml/types"
)

// func firstLetterIsLowercase(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}

// 	firstChar := rune(s[0])

// 	return unicode.IsLower(firstChar)
// }

func getTypeSpecFieldStructs(Expr ast.Expr) *[]types.StructField {
	var Structs []types.StructField

	switch Type := Expr.(type) {
	case *ast.Ident:
		Structs = append(Structs, types.StructField{
			Name: Type.Name,
		})

	case *ast.StructType:
		for _, field := range Type.Fields.List {
			for _, name := range field.Names {
				Structs = append(Structs, types.StructField{
					Name: name.Name,
				})
			}
		}
	case *ast.StarExpr:
		if ident, ok := Type.X.(*ast.Ident); ok {
			Structs = append(Structs, types.StructField{
				Name: ident.Name,
			})
		}
	case *ast.InterfaceType:
		for _, field := range Type.Methods.List {
			for _, name := range field.Names {
				Structs = append(Structs, types.StructField{
					Name: name.Name,
				})
			}
		}
	}

	return &Structs
}

func GetTypeSpecStruct(TypeSpec ast.TypeSpec) types.StructModel {
	TypeName, Fields, Methods := GetFieldType(TypeSpec.Type)

	Type := types.Type
	return types.StructModel{
		Header: types.StructHeader{
			Name:  TypeSpec.Name.Name,
			Title: TypeSpec.Name.Name + " " + TypeName,
		},

		Type:    &Type,
		Fields:  &Fields,
		Methods: &Methods,
	}
}
