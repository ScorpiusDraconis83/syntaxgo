package utils

import (
	"github.com/yyle88/printgo"
)

// NewPTX creates a new printgo PTX instance
// Returns a printgo PTX object that can be used to print Go data structures
//
// NewPTX 创建新的 printgo PTX 实例
// 返回可用于打印 Go 数据结构的 printgo PTX 对象
func NewPTX() *printgo.PTX {
	return printgo.NewPTX()
}
