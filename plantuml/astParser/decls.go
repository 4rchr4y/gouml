package astParser

import (
	"go/ast"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func GetDeclsStructs(Decls []ast.Decl) []types.StructModel {
	var Structs []types.StructModel

	for _, Decl := range Decls {
		switch d := Decl.(type) {
		case *ast.GenDecl:
			for _, Spec := range d.Specs {
				if importSpec, isImportSpec := Spec.(*ast.ImportSpec); isImportSpec {
					Structs = append(Structs, GetImportStruct(importSpec))
				}
				// if valueSpec, isValueSpec := Spec.(*ast.ValueSpec); isValueSpec {
				// 	Structs = append(Structs, GetValueSpecStruct(valueSpec))
				// }
				if typeSpec, isTypeSpec := Spec.(*ast.TypeSpec); isTypeSpec {
					Structs = append(Structs, GetTypeSpecStruct(*typeSpec))
				}
			}

		case *ast.FuncDecl:
			// FuncDecl struct {
			// 	Doc  *CommentGroup // associated documentation; or nil
			// 	Recv *FieldList    // receiver (methods); or nil (functions)
			// 	Name *Ident        // function/method name
			// 	Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
			// 	Body *BlockStmt    // function body; or nil for external (non-Go) function
			// }

		}
	}

	return Structs
}
