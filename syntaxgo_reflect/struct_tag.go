package syntaxgo_reflect

import "reflect"

// NewStructTag creates a reflect.StructTag from a string
// Converts a tag string into a reflect.StructTag type that can be queried
//
// NewStructTag 从字符串创建 reflect.StructTag
// 将标签字符串转换为可查询的 reflect.StructTag 类型
func NewStructTag(tag string) reflect.StructTag {
	return reflect.StructTag(tag)
}
