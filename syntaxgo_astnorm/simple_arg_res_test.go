package syntaxgo_astnorm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

// TestGetSimpleResElements tests extracting result elements from function return types
// Verifies GetSimpleResElements can parse return value information from AST
//
// TestGetSimpleResElements 测试从函数返回类型提取结果元素
// 验证 GetSimpleResElements 能够从 AST 解析返回值信息
func TestGetSimpleResElements(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	source := rese.A1(os.ReadFile(path))
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(source))
	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "GetSimpleResElements")
	require.NotNil(t, resFunc)

	elements := GetSimpleResElements(resFunc.Type.Results.List, source)
	for _, elem := range elements {
		t.Log(elem.Name, elem.Kind)
	}
}

// TestGetSimpleArgElements tests extracting argument elements from function parameters
// Verifies GetSimpleArgElements can parse parameter information from AST
//
// TestGetSimpleArgElements 测试从函数参数提取参数元素
// 验证 GetSimpleArgElements 能够从 AST 解析参数信息
func TestGetSimpleArgElements(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	source := rese.A1(os.ReadFile(path))
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(source))
	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "GetSimpleArgElements")
	require.NotNil(t, resFunc)

	elements := GetSimpleArgElements(resFunc.Type.Params.List, source)
	for _, elem := range elements {
		t.Log(elem.Name, elem.Kind)
	}
}
