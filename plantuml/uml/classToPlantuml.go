package uml

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func structModelToPlantumlSyntax(Struct types.StructModel) string {
	var sb strings.Builder

	// type StructModel struct {
	// 	Name       ident
	// 	Type       *structModelTypes
	// 	Visibility VisibilityKind
	// 	style      *structStyle
	// 	extends    string
	// 	Field      *[]structFieldAndMethod
	// 	Methods    *[]structFieldAndMethod
	// }

	// TODO я вообще блять не понимаю, как я должен из VisibilityKind получить value токена

	sb.WriteString("+")

	sb.WriteString(string(*Struct.Type))
	sb.WriteString(" ")

	if Struct.Name.Title == "" {
		sb.WriteString(Struct.Name.Name)
	} else {
		sb.WriteString("\"")
		sb.WriteString(Struct.Name.Title)
		sb.WriteString("\"")
		sb.WriteString(" as ")
		sb.WriteString(Struct.Name.Name)
	}

	sb.WriteString("{\n")

	for i, field := range *Struct.Field {
		if i > 0 {
			sb.WriteString("\n")
		}

		sb.WriteString("\t")
		sb.WriteString("+")
		sb.WriteString(string(field.Name))
	}

	sb.WriteString("\n}")

	return sb.String()
}

func AstToPlantuml(file *ast.File) {

	// name := file.Name

	// imports := types.GetImportsNames(file.Imports)
	var PlantumlSyntax strings.Builder

	for _, decl := range file.Decls {
		switch declType := decl.(type) {
		case *ast.FuncDecl:
			// fmt.Printf("Function: %s\n", declType.Name.Name)
		case *ast.GenDecl:
			if declType.Tok == token.IMPORT {
				continue
			}
			if declType.Tok == token.TYPE || declType.Tok == token.IDENT {
				Structs := types.GetStructFromDecl(declType)

				for _, Struct := range Structs {
					PlantumlSyntax.WriteString(structModelToPlantumlSyntax(Struct))
					PlantumlSyntax.WriteString("\n")
				}
			}
		}
	}
	fmt.Println(PlantumlSyntax.String())

}
