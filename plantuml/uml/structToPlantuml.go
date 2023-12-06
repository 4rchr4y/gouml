package uml

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/4rchr4y/gouml/plantuml/astParser"
	"github.com/4rchr4y/gouml/plantuml/types"
)

func structModelToPlantumlSyntax(Struct types.StructModel) string {
	var sb strings.Builder

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

	if Struct.Field != nil {
		for i, field := range *Struct.Field {
			if i > 0 {
				sb.WriteString("\n")
			}

			sb.WriteString("\t")
			sb.WriteString("+")
			sb.WriteString(string(field.Name))
		}

	}

	sb.WriteString("\n}\n")

	return sb.String()
}

func AstToPlantuml(file *ast.File) {
	var PlantumlSyntax strings.Builder

	// name := file.Name

	declsStructs := astParser.GetDeclsStructs(file.Decls)

	for _, decl := range declsStructs {
		PlantumlSyntax.WriteString(structModelToPlantumlSyntax(decl))
	}

	fmt.Println(PlantumlSyntax.String())

}
