// Demo2x shows type reflection and package path extraction
// Extract package info from Go types using reflection
// Generate qualified type names for code generation
//
// Demo2x 展示类型反射和包路径提取
// 使用反射从 Go 类型提取包信息
// 生成用于代码生成的限定类型名
package main

import (
	"fmt"
	"reflect"

	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// MyStruct is an example struct with basic fields
// Used to show type reflection and path extraction
//
// MyStruct 是带有基本字段的示例结构体
// 用于展示类型反射和路径提取
type MyStruct struct {
	Name string // User name field // 用户名字段
	Age  int    // User age field // 用户年龄字段
}

func main() {
	// Extract package path from type instance
	// Package path is the import path where type is defined
	//
	// 从类型实例提取包路径
	// 包路径是定义类型的导入路径
	pkgPath := syntaxgo_reflect.GetPkgPath(MyStruct{})
	fmt.Println("Package path:", pkgPath)

	// Generate qualified type usage code
	// Format: packageName.TypeName for use in code generation
	//
	// 生成限定的类型使用代码
	// 格式：packageName.TypeName 用于代码生成
	typeCode := syntaxgo_reflect.GenerateTypeUsageCode(reflect.TypeOf(MyStruct{}))
	fmt.Println("Type usage code:", typeCode)
}
