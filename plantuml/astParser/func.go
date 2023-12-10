package astParser

import (
	"strings"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func NewFunction(Function types.FunctionModel) types.StructModel {
	var title strings.Builder
	title.WriteString(Function.Name)
	title.WriteString("(")
	for i, param := range *Function.Params {
		if i > 0 {
			title.WriteString(", ")
		}
		title.WriteString(param.Name)
		title.WriteString(" ")
		title.WriteString(param.Type)
	}
	title.WriteString(")")

	Type := types.Import //TODO ПОЧЕМУ, БЛЯДЬ, Я НЕ МОГУ ПРОСТО ИСПОЛЬЗОВАТЬ ВНУТРИ СТРУКТУРЫ, СУКА

	return types.StructModel{
		Header: types.StructHeader{
			Name:  Function.Name,
			Title: title.String(),
		},
		Type:   &Type,
		Fields: &[]types.StructField{{Name: Function.ReturnType}},
	}
}
