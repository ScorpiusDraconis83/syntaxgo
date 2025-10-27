package syntaxgo_ast

import (
	"go/token"
	"testing"

	"github.com/yyle88/done"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

// TestGetPackageNameFromPath tests package name extraction from file path
// Verifies optimized parsing can get package name
//
// TestGetPackageNameFromPath 测试从文件路径提取包名
// 验证优化解析可以获取包名
func TestGetPackageNameFromPath(t *testing.T) {
	t.Log(GetPackageNameFromPath(runpath.Path()))
}

// TestGetPackageNameFromFile tests package name extraction from AST file node
// Verifies package name can be obtained from parsed AST structure
//
// TestGetPackageNameFromFile 测试从 AST 文件节点提取包名
// 验证可以从已解析的 AST 结构获取包名
func TestGetPackageNameFromFile(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	astBundle := done.P1(NewAstBundleV3(token.NewFileSet(), path))
	astFile := astBundle.file
	t.Log(GetPackageNameFromFile(astFile))
}
