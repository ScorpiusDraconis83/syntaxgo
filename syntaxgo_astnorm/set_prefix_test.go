package syntaxgo_astnorm

import (
	"go/ast"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

// TestNewPrefixedNameTypeElements tests creating elements with custom prefix
// Verifies NewPrefixedNameTypeElements can generate names with specified prefix
//
// TestNewPrefixedNameTypeElements 测试使用自定义前缀创建元素
// 验证 NewPrefixedNameTypeElements 能够使用指定前缀生成名称
func TestNewPrefixedNameTypeElements(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	source := rese.A1(os.ReadFile(path))
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(source))
	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "NewPrefixedNameTypeElements")
	require.NotNil(t, resFunc)

	elements := NewPrefixedNameTypeElements(
		resFunc.Type.Params,
		"abc88_",
		source,
		"pkg",
		map[string]ast.Expr{},
	)
	for _, elem := range elements {
		t.Log(elem.Name, elem.Kind)
	}
}
