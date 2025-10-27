package syntaxgo_astnode

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewNode tests Node creation and interface satisfaction
// Verifies that Node and NewNodeV1 create values that implement ast.Node interface
//
// TestNewNode 测试 Node 创建和接口满足
// 验证 Node 和 NewNodeV1 创建的值实现了 ast.Node 接口
func TestNewNode(t *testing.T) {
	var _ ast.Node = NewNode(1, 100)
	var _ ast.Node = NewNodeV1(NewNode(1, 100))
}

// TestNode_GetCode tests extracting code bytes from source using node positions
// Verifies GetCode can extract the correct byte content from source
//
// TestNode_GetCode 测试使用节点位置从源码提取代码字节
// 验证 GetCode 能够从源码提取正确的字节内容
func TestNode_GetCode(t *testing.T) {
	node := NewNode(1, 3)
	code := node.GetCode([]byte("abc"))
	t.Log(string(code))
	require.Equal(t, "ab", string(code))
}

// TestNode_GetText tests extracting code text from source using node positions
// Verifies GetText can extract the correct text string from source
//
// TestNode_GetText 测试使用节点位置从源码提取代码文本
// 验证 GetText 能够从源码提取正确的文本字符串
func TestNode_GetText(t *testing.T) {
	node := NewNode(1, 3)
	text := node.GetText([]byte("abc"))
	t.Log(text)
	require.Equal(t, "ab", text)
}
