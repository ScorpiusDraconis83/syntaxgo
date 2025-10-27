package syntaxgo_reflect

import "testing"

// TestGetPkgPaths tests extracting package paths from reflect types
// Verifies GetPkgPaths can get package paths from value types
//
// TestGetPkgPaths 测试从 reflect 类型提取包路径
// 验证 GetPkgPaths 能够从值类型获取包路径
func TestGetPkgPaths(t *testing.T) {
	objectsTypes := GetTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths(objectsTypes)
	t.Log(pkgPaths)
}

// TestGetPkgPathsFromPointers tests extracting package paths from pointer types
// Verifies GetPkgPaths can handle pointer types and extract correct paths
//
// TestGetPkgPathsFromPointers 测试从指针类型提取包路径
// 验证 GetPkgPaths 能够处理指针类型并提取正确路径
func TestGetPkgPathsFromPointers(t *testing.T) {
	objectsTypes := GetTypes([]any{&Example{}, &ExampleOneOne{}, &ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths(objectsTypes)
	t.Log(pkgPaths)
}

// TestGetPkgPathsToImportWithQuotes tests generating quoted import paths
// Verifies GetQuotedPackageImportPaths adds quotes for import statements
//
// TestGetPkgPathsToImportWithQuotes 测试生成带引号的导入路径
// 验证 GetQuotedPackageImportPaths 为导入语句添加引号
func TestGetPkgPathsToImportWithQuotes(t *testing.T) {
	objectsTypes := GetTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetQuotedPackageImportPaths(objectsTypes)
	t.Log(pkgPaths)
}
