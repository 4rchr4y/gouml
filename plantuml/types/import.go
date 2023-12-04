package types

import "go/ast"

func getImportName(ImportSpec *ast.ImportSpec) string {
	return ImportSpec.Path.Value
}

func GetImportsNames(Imports []*ast.ImportSpec) []string {
	var ImportsStrings []string

	for _, ImportSpec := range Imports {
		ImportsStrings = append(ImportsStrings, getImportName(ImportSpec))
	}

	return ImportsStrings
}
