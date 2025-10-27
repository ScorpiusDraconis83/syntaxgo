package syntaxgo_astnorm

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

func checkNiceFunction[A comparable, B comparable](a A, b A, x B) {
	must.Nice(a)
	must.Nice(b)
	must.Nice(x)
}

// TestGetFuncGenericTypeParamsMap tests extracting generic type parameters from function
// Verifies GetFuncGenericTypeParamsMap can parse generic type constraints from function declaration
//
// TestGetFuncGenericTypeParamsMap 测试从函数提取泛型类型参数
// 验证 GetFuncGenericTypeParamsMap 能够从函数声明解析泛型类型约束
func TestGetFuncGenericTypeParamsMap(t *testing.T) {
	checkNiceFunction("a", "b", 100)

	path := runpath.Path()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path))

	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "checkNiceFunction")
	require.NotNil(t, resFunc)

	nameMap := GetFuncGenericTypeParamsMap(resFunc)
	t.Log(nameMap) // output: map[A:comparable B:comparable]
}
