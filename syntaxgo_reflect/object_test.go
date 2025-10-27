package syntaxgo_reflect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetObject tests creating zero value objects with generics
// Verifies GetObject returns correct zero value for the specified type
//
// TestGetObject 测试使用泛型创建零值对象
// 验证 GetObject 为指定类型返回正确的零值
func TestGetObject(t *testing.T) {
	a := GetObject[int]()
	require.Equal(t, 0, a)
}

// TestNewObject tests creating pointer to zero value with generics
// Verifies NewObject returns pointer to correct zero value
//
// TestNewObject 测试使用泛型创建指向零值的指针
// 验证 NewObject 返回指向正确零值的指针
func TestNewObject(t *testing.T) {
	p := NewObject[int]()
	require.Equal(t, 0, *p)
}
