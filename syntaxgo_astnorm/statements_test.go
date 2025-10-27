package syntaxgo_astnorm

import "testing"

// TestStatementParts_MergeParts tests merging statement parts with commas
// Verifies MergeParts can join elements with comma and space separators
//
// TestStatementParts_MergeParts 测试用逗号合并语句部分
// 验证 MergeParts 能够用逗号和空格分隔符连接元素
func TestStatementParts_MergeParts(t *testing.T) {
	a := StatementParts{
		"n int", "s string", "v float64",
	}
	t.Log(a.MergeParts())
}

// TestStatementLines_MergeLines tests merging statement lines with newlines
// Verifies MergeLines can join elements with newline separators
//
// TestStatementLines_MergeLines 测试用换行符合并语句行
// 验证 MergeLines 能够用换行符分隔符连接元素
func TestStatementLines_MergeLines(t *testing.T) {
	a := StatementLines{
		"var n int",
		"var s string",
		"var v float64",
	}
	t.Log(a.MergeLines())
}
