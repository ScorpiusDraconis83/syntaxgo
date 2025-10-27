package syntaxgo_astnode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSdxEdx tests the SdxEdx function that converts node positions to indices
// Verifies correct conversion of AST node positions to zero-based indices
//
// TestSdxEdx 测试 SdxEdx 函数将节点位置转换为索引
// 验证 AST 节点位置正确转换为从零开始的索引
func TestSdxEdx(t *testing.T) {
	node := NewNode(1, 3)
	sdx, edx := SdxEdx(node)
	t.Log(sdx, edx)
	require.Equal(t, 0, sdx)
	require.Equal(t, 2, edx)
}

// TestGetCode tests extracting code bytes from source using node positions
// Verifies GetCode can extract the correct byte slice from source
//
// TestGetCode 测试使用节点位置从源码提取代码字节
// 验证 GetCode 能够从源码提取正确的字节切片
func TestGetCode(t *testing.T) {
	node := NewNode(1, 3)
	code := GetCode([]byte("abc"), node)
	t.Log(string(code))
	require.Equal(t, "ab", string(code))
}

// TestGetText tests extracting code text from source using node positions
// Verifies GetText can extract the correct text string from source
//
// TestGetText 测试使用节点位置从源码提取代码文本
// 验证 GetText 能够从源码提取正确的文本字符串
func TestGetText(t *testing.T) {
	node := NewNode(1, 3)
	text := GetText([]byte("abc"), node)
	t.Log(text)
	require.Equal(t, "ab", text)
}

// TestDeleteNodeCode tests removing node code from source
// Verifies DeleteNodeCode can remove specified node content from source
//
// TestDeleteNodeCode 测试从源码中删除节点代码
// 验证 DeleteNodeCode 能够从源码删除指定节点内容
func TestDeleteNodeCode(t *testing.T) {
	require.Equal(t, "a", string(DeleteNodeCode([]byte("abc"), NewNode(2, 4))))
	require.Equal(t, "ac", string(DeleteNodeCode([]byte("abc"), NewNode(2, 3))))
	require.Equal(t, "c", string(DeleteNodeCode([]byte("abc"), NewNode(1, 3))))
}

// TestChangeNodeCode tests replacing node code in source
// Verifies ChangeNodeCode can replace node content with new code
//
// TestChangeNodeCode 测试替换源码中的节点代码
// 验证 ChangeNodeCode 能够用新代码替换节点内容
func TestChangeNodeCode(t *testing.T) {
	require.Equal(t, "a123", string(ChangeNodeCode([]byte("abc"), NewNode(2, 4), []byte("123"))))
	require.Equal(t, "a88c", string(ChangeNodeCode([]byte("abc"), NewNode(2, 3), []byte("88"))))
	require.Equal(t, "666c", string(ChangeNodeCode([]byte("abc"), NewNode(1, 3), []byte("666"))))
}

// TestChangeNodeCodeSetSomeNewLines tests replacing node code with newlines
// Verifies code replacement with added newlines before and after new code
//
// TestChangeNodeCodeSetSomeNewLines 测试替换节点代码并添加换行
// 验证在新代码前后添加换行的代码替换
func TestChangeNodeCodeSetSomeNewLines(t *testing.T) {
	require.Equal(t, "a\n123\n", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(2, 4), []byte("123"), 1)))
	require.Equal(t, "a\n\n88\n\nc", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(2, 3), []byte("88"), 2)))
	require.Equal(t, "\n\n\n666\n\n\nc", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(1, 3), []byte("666"), 3)))
}
