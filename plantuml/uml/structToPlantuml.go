package uml

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/4rchr4y/gouml/plantuml/astParser"
	"github.com/4rchr4y/gouml/plantuml/types"
)

func structHeaderToPlantumlSyntax(Struct types.StructModel) string {
	var Header strings.Builder

	Header.WriteString("struct ")

	if Struct.Header.Title == "" {
		Header.WriteString(Struct.Header.Name)
	} else {
		Header.WriteString("\"")
		Header.WriteString(Struct.Header.Title)
		Header.WriteString("\"")
		Header.WriteString(" as ")
		Header.WriteString(Struct.Header.Name)
		Header.WriteString(" ")
	}

	if Struct.TitleStyle != nil {
		Header.WriteString("<<(")

		Header.WriteString(Struct.TitleStyle.Letter)
		Header.WriteString(",")
		Header.WriteString(Struct.TitleStyle.Color)

		Header.WriteString(")>> ")
	} else if Struct.Type != nil {
		switch *Struct.Type {
		case types.Import:
			Header.WriteString(" <<(I,SlateBlue)>> ")
		case types.Type:
			Header.WriteString(" <<(T,#fb432f)>> ")
		case types.Var:
			Header.WriteString(" <<(V,OrangeRed)>> ")
		case types.Function:
			Header.WriteString(" <<(F,HotPink)>> ")
		}
	}

	return Header.String()
}

func structModelToPlantumlSyntax(Struct types.StructModel) string {
	var sb strings.Builder

	// TODO я вообще блять не понимаю, как я должен из VisibilityKind получить value токена

	sb.WriteString(structHeaderToPlantumlSyntax(Struct))

	sb.WriteString("{\n")

	if Struct.Fields != nil {
		for i, field := range *Struct.Fields {
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
