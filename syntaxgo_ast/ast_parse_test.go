package syntaxgo_ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

// TestAstBundle_Print tests the Print method using the test file path
// Verifies AST structure can be printed to console output
//
// TestAstBundle_Print 使用测试文件路径测试 Print 方法
// 验证 AST 结构可以打印到控制台输出
func TestAstBundle_Print(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	astBundle := done.P1(NewAstBundleV3(token.NewFileSet(), path))
	done.Done(astBundle.Print())
}

// TestAstBundle_PrintCurrentFile tests printing AST of the current executing file
// Uses runtime path resolution to locate and parse the current Go source file
//
// TestAstBundle_PrintCurrentFile 测试打印当前执行文件的 AST
// 使用运行时路径解析来定位并解析当前 Go 源文件
func TestAstBundle_PrintCurrentFile(t *testing.T) {
	path := runpath.CurrentPath()
	t.Log(path)
	astBundle := rese.P1(NewAstBundleV4(path))
	done.Done(astBundle.Print())
}

// TestNewAstBundleV5 tests AstBundle creation with PackageClauseOnly mode
// Verifies parsing optimization when just package declaration is needed
//
// TestNewAstBundleV5 使用 PackageClauseOnly 模式测试 AstBundle 创建
// 验证当仅需要包声明时的解析优化
func TestNewAstBundleV5(t *testing.T) {
	astBundle, err := NewAstBundleV5(token.NewFileSet(), runtestpath.SrcPath(t), parser.PackageClauseOnly)
	require.NoError(t, err)
	done.Done(astBundle.Print())
	t.Log(string(done.VAE(astBundle.FormatSource()).Nice()))
}

// TestNewAstBundleV6 tests AstBundle creation with ImportsOnly mode
// Verifies parsing optimization when just import statements are needed
//
// TestNewAstBundleV6 使用 ImportsOnly 模式测试 AstBundle 创建
// 验证当仅需要导入语句时的解析优化
func TestNewAstBundleV6(t *testing.T) {
	astBundle, err := NewAstBundleV6(runpath.Path(), parser.ImportsOnly)
	require.NoError(t, err)
	done.Done(astBundle.Print())
	t.Log(string(done.VAE(astBundle.FormatSource()).Nice()))
}

// TestAstBundle_FormatSource tests formatting AST back to Go source code
// Verifies import injection and source code generation from AST
//
// TestAstBundle_FormatSource 测试将 AST 格式化回 Go 源代码
// 验证导入注入和从 AST 生成源代码
func TestAstBundle_FormatSource(t *testing.T) {
	const code = `package main
	//这是main函数的注释
	func main() {
		fmt.Println("abc") //随便打印个字符串
		fmt.Println(time.Now()) //随便打印当前时间
	}
`
	astBundle, err := NewAstBundleV2(token.NewFileSet(), []byte(code))
	require.NoError(t, err)

	added := astBundle.AddImport("fmt")
	require.True(t, added)

	added = astBundle.AddImport("time")
	require.True(t, added)

	newSrc, err := astBundle.FormatSource()
	require.NoError(t, err)
	t.Log(string(newSrc))
}

// TestAstBundle_SerializeAst tests AST serialization with embed directives
// Verifies named import handling and directive preservation in AST
//
// TestAstBundle_SerializeAst 测试带有 embed 指令的 AST 序列化
// 验证命名导入处理和 AST 中的指令保留
func TestAstBundle_SerializeAst(t *testing.T) {
	const code = `package main

	//go:embed hello.txt
	var s string

	//这是main函数的注释
	func main() {
		fmt.Println(s) //打印整个文件内容
	}
`
	astBundle, err := NewAstBundleV2(token.NewFileSet(), []byte(code))
	require.NoError(t, err)

	added := astBundle.AddImport("fmt")
	require.True(t, added)

	added = astBundle.AddNamedImport("_", "embed")
	require.True(t, added)

	newSrc, err := astBundle.SerializeAst()
	require.NoError(t, err)
	t.Log(string(newSrc))
}

// TestAstBundle_GetPackageName tests package name extraction from AST
// Verifies optimized parsing with PackageClauseOnly mode
//
// TestAstBundle_GetPackageName 测试从 AST 提取包名
// 验证使用 PackageClauseOnly 模式的优化解析
func TestAstBundle_GetPackageName(t *testing.T) {
	astBundle := rese.P1(NewAstBundleV6(runpath.Path(), parser.PackageClauseOnly))
	packageName := astBundle.GetPackageName()
	t.Log(packageName)
}

// TestAstBundle_CheckTokenType tests token type identification in AST declarations
// Verifies TYPE token detection and type specification extraction
//
// TestAstBundle_CheckTokenType 测试 AST 声明中的 token 类型识别
// 验证 TYPE token 检测和类型规范提取
func TestAstBundle_CheckTokenType(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	astBundle := done.P1(NewAstBundleV3(token.NewFileSet(), path))
	astFile, _ := astBundle.GetBundle()

	for _, declaration := range astFile.Decls {
		genericDeclaration, ok := declaration.(*ast.GenDecl)
		if !ok {
			continue
		}
		t.Log("type_enum:", genericDeclaration.Tok)
		if genericDeclaration.Tok == token.TYPE {
			for _, spec := range genericDeclaration.Specs {
				if typeDeclaration, ok := spec.(*ast.TypeSpec); ok {
					t.Log("type_name:", typeDeclaration.Name.Name)
				}
			}
		}
	}
}
