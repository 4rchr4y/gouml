package astParser

import (
	"go/ast"
	"strings"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func getLastWord(input string) string {
	parts := strings.Split(input, "/")

	if len(parts) == 0 {
		return ""
	}

	return strings.Replace(parts[len(parts)-1], "\"", "", -1)
}

func GetImportName(ImportSpec *ast.ImportSpec) string {
	return getLastWord(ImportSpec.Path.Value)
}

func GetImportStruct(ImportSpec *ast.ImportSpec) types.StructModel {
	Type := types.Interface
	return types.StructModel{
		Name: types.Ident{Name: GetImportName(ImportSpec)},
		Type: &Type,
	}
}

func GetImportsStructs(Imports []*ast.ImportSpec) []types.StructModel {
	var ImportsStructs []types.StructModel
	Type := types.Interface

	for _, ImportSpec := range Imports {
		name := GetImportName(ImportSpec)
		ImportsStructs = append(ImportsStructs, types.StructModel{
			Name: types.Ident{
				Name: name,
			},
			Type: &Type,
		})
	}

	return ImportsStructs
}
