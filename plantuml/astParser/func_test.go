package astParser

import (
	"fmt"
	"testing"

	"github.com/4rchr4y/gouml/plantuml/types"
)

func TestNewFunction(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		model := types.FunctionModel{
			Name: "SetToken",
			Params: &[]types.FunctionParamModel{
				{Name: "foo", Type: "int"},
			},
			ReturnType: "string",
		}

		res := NewFunction(model)

		fmt.Println(res)
		// assert.NotNil(t, weakRef, "unexpected: weakRef is nil")
	})
}
