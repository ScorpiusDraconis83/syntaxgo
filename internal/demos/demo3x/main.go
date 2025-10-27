// Demo3x shows AST search and code element finding
// Search functions and types in parsed AST
// Navigate struct fields and extract declarations
//
// Demo3x 展示 AST 搜索和代码元素查找
// 在已解析的 AST 中搜索函数和类型
// 遍历结构体字段并提取声明
package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/yyle88/must"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

func main() {
	// Parse Go source code from string
	// Create token FileSet to track positions
	//
	// 从字符串解析 Go 源代码
	// 创建 token FileSet 来跟踪位置
	fset := token.NewFileSet()
	src := `package example

func HelloWorld() string {
	return "Hello, World!"
}

type Person struct {
	Name string
	Age  int
}
`
	astFile, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Search function declaration in AST
	// Find function node with specified name
	//
	// 在 AST 中搜索函数声明
	// 查找具有指定名称的函数节点
	astFunc := syntaxgo_search.FindFunctionByName(astFile, "HelloWorld")
	if astFunc != nil {
		fmt.Println("Found function:", astFunc.Name.Name)
	}

	// Search struct type definition in AST
	// Locate struct node and extract its fields
	//
	// 在 AST 中搜索结构体类型定义
	// 定位结构体节点并提取其字段
	structType, ok := syntaxgo_search.FindStructTypeByName(astFile, "Person")
	must.OK(ok)

	// Traverse struct fields and print field names
	// Each field can have multiple names in one line
	//
	// 遍历结构体字段并打印字段名
	// 每个字段可以在一行中有多个名称
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			fmt.Println("Person Field:", fieldName.Name)
		}
	}
}
