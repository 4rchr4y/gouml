package astParser

import (
	"go/ast"
	"strings"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func GetValueSpecStruct(Var *ast.ValueSpec) types.StructModel {
	var Struct types.StructModel

	return Struct
}
func NewVar(Var types.VarModel) types.StructModel {
	var title strings.Builder
	title.WriteString(Var.Name)
	title.WriteString(" ")
	title.WriteString(Var.Type)

	// TODO fileds(var value)
	Type := types.Var
	return types.StructModel{
		Header: types.StructHeader{
			Name:  Var.Name,
			Title: title.String(),
		},
		Type: &Type,
	}
}
