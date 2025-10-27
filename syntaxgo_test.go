package syntaxgo_test

import (
	"testing"

	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
)

// TestCurrentPackageName tests getting the current package name at runtime
// Verifies CurrentPackageName function can detect the calling package name
//
// TestCurrentPackageName 测试运行时获取当前包名
// 验证 CurrentPackageName 函数能够检测调用者的包名
func TestCurrentPackageName(t *testing.T) {
	t.Log(syntaxgo.CurrentPackageName())
}

// TestGetPkgName tests extracting package name from a Go source file path
// Verifies GetPkgName function can parse package declaration from file
//
// TestGetPkgName 测试从 Go 源文件路径提取包名
// 验证 GetPkgName 函数能够从文件解析包声明
func TestGetPkgName(t *testing.T) {
	t.Log(syntaxgo.GetPkgName(runtestpath.SrcPath(t)))
}
