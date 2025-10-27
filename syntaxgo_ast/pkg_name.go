package syntaxgo_ast

import (
	"go/ast"
	"go/parser"

	"github.com/yyle88/rese"
)

// GetPackageNameFromPath extracts package name from a Go source file at the given path
// Uses optimized parsing with PackageClauseOnly mode to read just the package statement
// Returns the package name as a string without reading the entire file content
//
// GetPackageNameFromPath 从给定路径的 Go 源文件中提取包名
// 使用优化的 PackageClauseOnly 模式解析，仅读取包声明语句
// 返回包名字符串，无需读取整个文件内容
func GetPackageNameFromPath(path string) string {
	return rese.P1(NewAstBundleV6(path, parser.PackageClauseOnly)).GetPackageName()
}

// GetPackageNameFromFile extracts package name from an AST file node
// Accesses the Name field of the AST file to obtain the package name
// Returns the package name as a string from the parsed AST structure
//
// GetPackageNameFromFile 从 AST 文件节点中提取包名
// 访问 AST 文件的 Name 字段以获取包名称
// 从已解析的 AST 结构返回包名字符串
func GetPackageNameFromFile(astFile *ast.File) (packageName string) {
	return astFile.Name.Name
}
