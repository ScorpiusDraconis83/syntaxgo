package syntaxgo_astnorm

import "strings"

// StatementParts represents a list of strings separated by commas
// Used in parameter lists, return value lists, and function arguments
//
// StatementParts 表示以逗号分隔的字符串列表
// 用于参数列表、返回值列表和函数参数
type StatementParts []string

// MergeParts joins the elements in StatementParts with a comma and a space, and returns the resulting string.
// MergeParts 将 StatementParts 中的元素用逗号和空格连接，并返回结果字符串。
func (stmts StatementParts) MergeParts() string {
	return strings.Join(stmts, ", ")
}

// StatementLines represents a list of strings separated by newlines
// Used in assignment statements, return statements, and function calls
//
// StatementLines 表示以换行符分隔的字符串列表
// 用于赋值语句、返回语句和函数调用
type StatementLines []string

// MergeLines joins the elements in StatementLines with a newline character, and returns the resulting string.
// MergeLines 将 StatementLines 中的元素用换行符连接，并返回结果字符串。
func (stmts StatementLines) MergeLines() string {
	return strings.Join(stmts, "\n")
}
