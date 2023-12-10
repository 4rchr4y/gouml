package astParser

import (
	"go/ast"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func GetFieldType(Expr ast.Expr) (string, []types.StructField, []types.StructMethod) {
	switch node := Expr.(type) {
	case *ast.Ident:
		return getIdentType(node)
	case *ast.ArrayType:
		return getArrayType(node)
	case *ast.SelectorExpr:
		return getSelectorExprType(node)
	case *ast.MapType:
		return getMapType(node), nil, nil
	case *ast.StarExpr:
		return getStarExprType(node)
	case *ast.ChanType:
		return getChanType(node)
	case *ast.StructType:
		return getStructType(node)
	case *ast.InterfaceType:
		return getInterfaceType(node), nil, nil
	case *ast.FuncType:
		return getFuncType(node), nil, nil
	case *ast.Ellipsis:
		return getEllipsisType(node)
	}

	return "", nil, nil
}

func getIdentType(Ident *ast.Ident) (string, []types.StructField, []types.StructMethod) {
	return Ident.Name, nil, nil
}

func getArrayType(ArrayType *ast.ArrayType) (string, []types.StructField, []types.StructMethod) {
	StructType, StructFields, StructMethods := GetFieldType(ArrayType.Elt)
	return "[]" + StructType, StructFields, StructMethods
}

func getStarExprType(StarExpr *ast.StarExpr) (string, []types.StructField, []types.StructMethod) {
	StructType, StructFields, StructMethods := GetFieldType(StarExpr.X)
	return "*" + StructType, StructFields, StructMethods
}

func getChanType(ChanType *ast.ChanType) (string, []types.StructField, []types.StructMethod) {
	StructType, StructFields, StructMethods := GetFieldType(ChanType.Value)
	return "<font color=blue>chan</font> " + StructType, StructFields, StructMethods
}

func getEllipsisType(Ellipsis *ast.Ellipsis) (string, []types.StructField, []types.StructMethod) {
	StructType, StructFields, StructMethods := GetFieldType(Ellipsis.Elt)
	return "..." + StructType, StructFields, StructMethods
}

// ----------------------------------------

func getStructType(Struct *ast.StructType) (string, []types.StructField, []types.StructMethod) {
	fields := []types.StructField{}

	for _, field := range Struct.Fields.List {
		fieldType, _, _ := GetFieldType(field.Type)
		fields = append(fields, types.StructField{
			Name: fieldType,
		})
	}

	return "<font color=blue>struct</font> ", fields, nil
}

func getInterfaceType(InterfaceType *ast.InterfaceType) string {
	return ""
}

func getSelectorExprType(SelectorExpr *ast.SelectorExpr) (string, []types.StructField, []types.StructMethod) {
	return GetFieldType(SelectorExpr.X)
}

func getMapType(MapType *ast.MapType) string {
	return ""
}

func getFuncType(FuncType *ast.FuncType) string {
	return ""
}
