package syntaxgo

import (
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

// CurrentPackageName returns the package name at the calling site
// Uses runtime stack inspection to skip one frame and locate the calling source file
// Then parses the package declaration to extract the package name
//
// CurrentPackageName 返回调用者位置的包名
// 使用运行时栈检查跳过一层来定位调用源文件
// 然后解析包声明以提取包名
func CurrentPackageName() string {
	return syntaxgo_ast.GetPackageNameFromPath(runpath.Skip(1))
}

// GetPkgName returns the package name from a Go source file path
// Parses the package declaration at the beginning of the file to extract the name
// Works with any valid Go source file containing a package statement
//
// GetPkgName 从 Go 源文件路径返回包名
// 解析文件开头的包声明来提取名称
// 适用于任何包含 package 语句的有效 Go 源文件
func GetPkgName(path string) string {
	return syntaxgo_ast.GetPackageNameFromPath(path)
}
