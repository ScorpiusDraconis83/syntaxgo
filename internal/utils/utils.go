// Package utils provides internal helper functions for syntaxgo
// These utilities serve the main syntaxgo package and are not exported
// Functions here handle common operations like string manipulation and file checks
//
// Design Decision: Keep utils internal to avoid circular dependencies
// Since syntaxgo is a base package, external tool packages might import it
// Extracting these utils to a separate package could create circular references
// While project-level circular deps compile, they complicate version management
//
// utils 包为 syntaxgo 提供内部辅助函数
// 这些工具服务于主 syntaxgo 包且不对外导出
// 此处函数处理常见操作，如字符串操作和文件检查
//
// 设计决策：保持 utils 为内部包以避免循环依赖
// 由于 syntaxgo 是基础包，外部工具包可能会引用它
// 将这些工具提取到独立包可能造成循环引用
// 虽然项目级循环依赖能编译，但会使版本管理复杂化
package utils

import (
	"os"
	"path/filepath"
	"unicode"
)

// SetPrefix2Strings adds a prefix to each string in the slice
// Returns a new slice with each string prefixed
//
// SetPrefix2Strings 为切片中的每个字符串添加前缀
// 返回包含每个添加前缀字符串的新切片
func SetPrefix2Strings(prefix string, a []string) (results []string) {
	results = make([]string, 0, len(a))
	for _, v := range a {
		results = append(results, prefix+v)
	}
	return results
}

// SafeMerge combines multiple slices into a single slice
// Pre-allocates space based on combined length to optimize performance
//
// SafeMerge 将多个切片合并为单个切片
// 根据组合长度预分配空间以优化性能
func SafeMerge[V any](a ...[]V) (res []V) {
	res = make([]V, 0, SumLength(a...))
	for _, slice := range a {
		res = append(res, slice...)
	}
	return res
}

// SumLength calculates the combined length of slices
// Returns the sum of lengths across input slices
//
// SumLength 计算所有切片的组合长度
// 返回输入切片的长度总和
func SumLength[V any](a ...[]V) (n int) {
	for _, slice := range a {
		n += len(slice)
	}
	return n
}

// C0IsUppercase checks if the first rune in the string is uppercase
// Returns false when the string has no content
//
// C0IsUppercase 检查字符串的第一个字符是否为大写
// 当字符串没有内容时返回 false
func C0IsUppercase(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 {
		return unicode.IsUpper(runes[0])
	}
	return false
}

// SetDoubleQuotes wraps a string with double quotes
// Returns the string surrounded with double quote characters
//
// SetDoubleQuotes 用双引号包装字符串
// 返回用双引号字符包围的字符串
func SetDoubleQuotes(s string) string {
	return "\"" + s + "\""
}

// IsGoSourceFile checks if the file info represents a Go source file
// Returns true if the file has .go extension and is not a DIR
//
// IsGoSourceFile 检查文件信息是否代表 Go 源文件
// 如果文件具有 .go 扩展名且不是 DIR 则返回 true
func IsGoSourceFile(info os.FileInfo) bool {
	return (!info.IsDir()) && (filepath.Ext(info.Name()) == ".go")
}
