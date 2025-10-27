// Demo1x shows basic AST parsing and manipulation
// Parse Go source file into AST and extract package info
// Print AST structure for inspection and debugging
//
// Demo1x 展示基础的 AST 解析和操作
// 将 Go 源文件解析为 AST 并提取包信息
// 打印 AST 结构用于检查和调试
package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func main() {
	// Parse Go source file from current path
	// Create AST bundle for analysis and manipulation
	//
	// 从当前路径解析 Go 源文件
	// 创建 AST bundle 用于分析和操作
	astBundle, err := syntaxgo_ast.NewAstBundleV4(runpath.Current())
	if err != nil {
		panic(err)
	}

	// Extract package name from AST
	// Package name is the first identifier in package clause
	//
	// 从 AST 提取包名
	// 包名是 package 子句中的第一个标识符
	pkgName := astBundle.GetPackageName()
	fmt.Println("Package name:", pkgName)

	// Print complete AST structure to console
	// Show all nodes including declarations and statements
	//
	// 打印完整的 AST 结构到控制台
	// 显示所有节点包括声明和语句
	must.Done(astBundle.Print())
}
