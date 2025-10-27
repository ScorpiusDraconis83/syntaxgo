package syntaxgo_reflect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

type Example struct{}

func (x *Example) methodFunction() {}

// TestGetPkgPath tests extracting package path and name from various types
// Verifies GetPkgPath and GetPkgName work with values, pointers, and methods
//
// TestGetPkgPath 测试从各种类型提取包路径和名称
// 验证 GetPkgPath 和 GetPkgName 能够处理值、指针和方法
func TestGetPkgPath(t *testing.T) {
	t.Log(GetPkgPath(Example{}))
	t.Log(GetPkgName(Example{}))

	t.Log(GetPkgPath(&Example{}))
	t.Log(GetPkgName(&Example{}))

	a := &Example{}
	t.Log(GetPkgPath(a.methodFunction))
	t.Log(GetPkgName(a.methodFunction))

	b := Example{}
	t.Log(GetPkgPath(b.methodFunction))
	t.Log(GetPkgName(b.methodFunction))
}

func commonFunction() {}

// TestGetPkgPath1 tests extracting package info from function values
// Verifies GetPkgPath and GetPkgName work with function types
//
// TestGetPkgPath1 测试从函数值提取包信息
// 验证 GetPkgPath 和 GetPkgName 能够处理函数类型
func TestGetPkgPath1(t *testing.T) {
	t.Log(GetPkgPath(commonFunction))
	t.Log(GetPkgName(commonFunction))
}

type ExampleInterface interface {
	methodFunction()
}

type ExampleOneOne struct{}

func (e *ExampleOneOne) methodFunction() {}

// TestGetPkgPath2 tests extracting package info from interface with pointer receiver
// Verifies GetPkgPath and GetPkgName handle interface values with pointer implementation
//
// TestGetPkgPath2 测试从带指针接收者的接口提取包信息
// 验证 GetPkgPath 和 GetPkgName 能够处理指针实现的接口值
func TestGetPkgPath2(t *testing.T) {
	var a ExampleInterface = &ExampleOneOne{}
	t.Log(GetPkgPath(a))
	t.Log(GetPkgName(a))
}

type ExampleTwoTwo struct{}

func (e ExampleTwoTwo) methodFunction() {}

// TestGetPkgPath3 tests extracting package info from interface with value receiver
// Verifies GetPkgPath and GetPkgName handle interface values with value implementation
//
// TestGetPkgPath3 测试从带值接收者的接口提取包信息
// 验证 GetPkgPath 和 GetPkgName 能够处理值实现的接口值
func TestGetPkgPath3(t *testing.T) {
	var a ExampleInterface = ExampleTwoTwo{}
	t.Log(GetPkgPath(a))
	t.Log(GetPkgName(a))
}

// TestGetTypeName tests extracting type name from value
// Verifies GetTypeName returns correct type name string
//
// TestGetTypeName 测试从值提取类型名称
// 验证 GetTypeName 返回正确的类型名称字符串
func TestGetTypeName(t *testing.T) {
	typeName := GetTypeName(Example{})
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

// TestGetTypeNameV2 tests extracting type name using generics
// Verifies GetTypeNameV2 returns correct type name via generic type parameter
//
// TestGetTypeNameV2 测试使用泛型提取类型名称
// 验证 GetTypeNameV2 通过泛型类型参数返回正确的类型名称
func TestGetTypeNameV2(t *testing.T) {
	typeName := GetTypeNameV2[Example]()
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

// TestGetTypeNameV3 tests extracting type name from pointer value
// Verifies GetTypeNameV3 extracts base type name from pointer types
//
// TestGetTypeNameV3 测试从指针值提取类型名称
// 验证 GetTypeNameV3 从指针类型提取基础类型名称
func TestGetTypeNameV3(t *testing.T) {
	typeName := GetTypeNameV3(&Example{})
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

// TestGetTypeNameV4 tests extracting type name from pointer with auto detection
// Verifies GetTypeNameV4 handles pointer type name extraction
//
// TestGetTypeNameV4 测试从指针自动检测提取类型名称
// 验证 GetTypeNameV4 处理指针类型名称提取
func TestGetTypeNameV4(t *testing.T) {
	typeName := GetTypeNameV4(&Example{})
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

// TestGetPkgPathV2 tests extracting package info using generics
// Verifies GetPkgPathV2 and GetPkgNameV2 work with generic type parameters
// Note: Pointer types are not supported and return empty strings
//
// TestGetPkgPathV2 测试使用泛型提取包信息
// 验证 GetPkgPathV2 和 GetPkgNameV2 能够使用泛型类型参数
// 注意：不支持指针类型，会返回空字符串
func TestGetPkgPathV2(t *testing.T) {
	t.Log(GetPkgPathV2[Example]())
	t.Log(GetPkgNameV2[Example]())

	t.Log(GetPkgPathV2[*Example]()) //这个是不行的，目前不支持往里面传指针类型，但并不会panic而是返回空白
	t.Log(GetPkgNameV2[*Example]()) //这个是不行的，目前不支持往里面传指针类型，但并不会panic而是返回空白
}

// TestGetPkgNameV3 tests extracting package info from values and pointers
// Verifies GetPkgPathV3 and GetPkgNameV3 handle both value and pointer types
//
// TestGetPkgNameV3 测试从值和指针提取包信息
// 验证 GetPkgPathV3 和 GetPkgNameV3 能够处理值类型和指针类型
func TestGetPkgNameV3(t *testing.T) {
	t.Log(GetPkgPathV3(Example{}))
	t.Log(GetPkgNameV3(Example{}))

	t.Log(GetPkgPathV3(&Example{}))
	t.Log(GetPkgNameV3(&Example{}))
}

// TestGetPkgNameV4 tests extracting package info from pointer types
// Verifies GetPkgPathV4 and GetPkgNameV4 handle pointer type introspection
//
// TestGetPkgNameV4 测试从指针类型提取包信息
// 验证 GetPkgPathV4 和 GetPkgNameV4 能够处理指针类型内省
func TestGetPkgNameV4(t *testing.T) {
	t.Log(GetPkgPathV4(&Example{}))
	t.Log(GetPkgNameV4(&Example{}))
}
