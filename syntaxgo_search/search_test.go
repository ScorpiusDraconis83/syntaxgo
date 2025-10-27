package syntaxgo_search

import (
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

// TestAstBundle_Print_Search tests printing AST structure of search module
// Verifies AST parsing and printing for the current test source file
//
// TestAstBundle_Print_Search 测试打印搜索模块的 AST 结构
// 验证当前测试源文件的 AST 解析和打印
func TestAstBundle_Print_Search(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle, err := syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path)
	require.NoError(t, err)
	done.Done(astBundle.Print())
}

// TestFindFunctionByName tests finding function by name in AST
// Verifies FindFunctionByName can locate and extract function documentation
//
// TestFindFunctionByName 测试在 AST 中按名称查找函数
// 验证 FindFunctionByName 能够定位并提取函数文档
func TestFindFunctionByName(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path))

	astFile, _ := astBundle.GetBundle()

	astFunc := FindFunctionByName(astFile, "FindFunctionByName")
	if astFunc == nil {
		return
	}
	for k, s := range astFunc.Doc.List {
		fmt.Println("-----", k, "-----")
		fmt.Println(s.Text)
		fmt.Println("-----", "-", "-----")
	}
}

type Example struct {
	Name string
}

type Examples []*Example

// TestFindArrayTypeByName tests finding array type declarations in AST
// Verifies FindArrayTypeByName can locate custom array type definitions
//
// TestFindArrayTypeByName 测试在 AST 中查找数组类型声明
// 验证 FindArrayTypeByName 能够定位自定义数组类型定义
func TestFindArrayTypeByName(t *testing.T) {
	examples := make(Examples, 0, 3)
	examples = append(examples, &Example{Name: "a"})
	examples = append(examples, &Example{Name: "b"})
	examples = append(examples, &Example{Name: "c"})
	t.Log(examples)

	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astFunc := FindArrayTypeByName(astFile, "Examples")
	require.NotNil(t, astFunc)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astFunc)))
}

// TestFindStructTypeByName tests finding struct type by name in AST
// Verifies FindStructTypeByName can locate struct definitions
//
// TestFindStructTypeByName 测试在 AST 中按名称查找结构体类型
// 验证 FindStructTypeByName 能够定位结构体定义
func TestFindStructTypeByName(t *testing.T) {
	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astStruct, ok := FindStructTypeByName(astFile, "Example")
	require.True(t, ok)
	require.NotNil(t, astStruct)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astStruct)))
}

// TestFindStructDeclarationByName tests finding complete struct declarations
// Verifies FindStructDeclarationByName extracts entire struct declaration
//
// TestFindStructDeclarationByName 测试查找完整的结构体声明
// 验证 FindStructDeclarationByName 能够提取完整的结构体声明
func TestFindStructDeclarationByName(t *testing.T) {
	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astStruct, ok := FindStructDeclarationByName(astFile, "Example")
	require.True(t, ok)
	require.NotNil(t, astStruct)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astStruct)))
}

// TestFindFunctions tests finding all functions in Go source files
// Verifies FindFunctions can extract function declarations and documentation
//
// TestFindFunctions 测试在 Go 源文件中查找所有函数
// 验证 FindFunctions 能够提取函数声明和文档
func TestFindFunctions(t *testing.T) {
	root := runpath.PARENT.Path()

	require.NoError(t, filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !utils.IsGoSourceFile(info) {
			return nil
		}
		if strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}

		astBundle := rese.P1(syntaxgo_ast.NewAstBundleV4(path))

		astFile, _ := astBundle.GetBundle()

		astFunctions := FindFunctions(astFile)

		for _, astFunction := range astFunctions {
			t.Log(astFunction.Name.Name, "//", GetFunctionComment(astFunction))
		}

		return nil
	}))
}
