package types

type FunctionModel struct {
	Name       string
	Params     *[]FunctionParamModel
	ReturnType string
}

type FunctionParamModel struct {
	Name string
	Type string
}
