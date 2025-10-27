[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/syntaxgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/syntaxgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/syntaxgo)](https://pkg.go.dev/github.com/yyle88/syntaxgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/syntaxgo/master.svg)](https://coveralls.io/github/yyle88/syntaxgo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/syntaxgo.svg)](https://github.com/yyle88/syntaxgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/syntaxgo)](https://goreportcard.com/report/github.com/yyle88/syntaxgo)

# syntaxgo

**syntaxgo** 是基于 Go 的 `go/ast` 抽象语法树和 `reflect` 反射包构建的 Go 工具包。

**syntaxgo** 旨在简化代码分析和自动生成。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 主要功能

🎯 **AST 解析**：将 Go 源代码解析为抽象语法树，支持多种解析模式
⚡ **类型反射**：提取包路径、类型信息和生成使用代码
🔍 **代码搜索**：在 AST 中查找函数、类型和结构体字段
📝 **标签操作**：提取和操作结构体字段标签
🛠️ **代码生成**：生成函数参数、返回值和变量定义
📦 **导入管理**：自动注入缺失的导入和管理包路径

## 安装

```bash
go get github.com/yyle88/syntaxgo
```

## 使用方法

### 基础 AST 解析

```go
package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func main() {
	// 从当前路径解析 Go 源文件
	astBundle, err := syntaxgo_ast.NewAstBundleV4(runpath.Current())
	if err != nil {
		panic(err)
	}

	// 从 AST 提取包名
	pkgName := astBundle.GetPackageName()
	fmt.Println("Package name:", pkgName)

	// 打印完整的 AST 结构到控制台
	must.Done(astBundle.Print())
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### 类型反射

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type MyStruct struct {
	Name string // 用户名字段
	Age  int    // 用户年龄字段
}

func main() {
	// 从类型实例提取包路径
	pkgPath := syntaxgo_reflect.GetPkgPath(MyStruct{})
	fmt.Println("Package path:", pkgPath)

	// 生成限定的类型使用代码
	typeCode := syntaxgo_reflect.GenerateTypeUsageCode(reflect.TypeOf(MyStruct{}))
	fmt.Println("Type usage code:", typeCode)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### AST 搜索

```go
package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/yyle88/must"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

func main() {
	// 从字符串解析 Go 源代码
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

	// 在 AST 中搜索函数声明
	astFunc := syntaxgo_search.FindFunctionByName(astFile, "HelloWorld")
	if astFunc != nil {
		fmt.Println("Found function:", astFunc.Name.Name)
	}

	// 在 AST 中搜索结构体类型定义
	structType, ok := syntaxgo_search.FindStructTypeByName(astFile, "Person")
	must.OK(ok)

	// 遍历结构体字段并打印字段名
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			fmt.Println("Person Field:", fieldName.Name)
		}
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

### 标签操作

```go
package main

import (
	"fmt"

	"github.com/yyle88/syntaxgo/syntaxgo_tag"
)

func main() {
	// 定义带有 GORM 和 JSON 注解的结构体字段标签
	tag := `gorm:"column:id;type:bigint" json:"id"`

	// 从完整标签字符串提取 GORM 标签值
	gormTag := syntaxgo_tag.ExtractTagValue(tag, "gorm")
	fmt.Println("GORM tag:", gormTag)

	// 从 GORM 标签提取 column 字段值
	column := syntaxgo_tag.ExtractTagField(gormTag, "column", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column name:", column)

	// 从 GORM 标签提取 type 字段值
	typeValue := syntaxgo_tag.ExtractTagField(gormTag, "type", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column type:", typeValue)
}
```

⬆️ **Source:** [Source](internal/demos/demo4x/main.go)

## 模块说明

### syntaxgo_ast - AST 解析与管理

**核心函数：**
- `NewAstBundleV1/V2/V3/V4/V5/V6` - 使用不同输入模式解析 Go 源文件（字节、路径、FileSet）
- `FormatSource` - 将 AST 格式化为规范的 Go 源代码
- `SerializeAst` - 将 AST 序列化为文本表达形式
- `GetPackageName` - 从 AST 提取包名
- `AddImport/DeleteImport` - 管理单个导入路径
- `AddNamedImport/DeleteNamedImport` - 处理别名导入（如 `_ "embed"`）
- `InjectImports` - 自动向源代码注入缺失的导入
- `CreateImports` - 从包路径生成导入块

**使用场景：**
- 解析源文件并保留注释
- 自动修复生成代码中缺失的导入
- 扫描项目时提取包信息
- AST 操作后格式化代码

### syntaxgo_astnode - AST 节点操作

**核心函数：**
- `NewNode/NewNodeV1/NewNodeV2` - 创建带位置信息的节点
- `GetCode/GetText` - 从 AST 节点提取代码内容
- `SdxEdx` - 获取节点的起始/结束索引
- `DeleteNodeCode` - 从源码删除节点代码
- `ChangeNodeCode` - 用新代码替换节点内容
- `ChangeNodeCodeSetSomeNewLines` - 替换并添加换行

**使用场景：**
- 从 AST 提取函数体
- 替换结构体字段定义
- 删除未使用的代码段
- 在特定位置插入新代码

### syntaxgo_astnorm - 函数签名处理

**核心函数：**
- `NewNameTypeElements` - 从 AST 字段列表提取参数/返回值
- `ExtractNameTypeElements` - 使用自定义命名处理字段列表
- `Names/Kinds` - 获取参数名称和类型的独立列表
- `FormatNamesWithKinds` - 格式化为 "名称 类型" 对
- `GenerateFunctionParams` - 创建调用函数时的参数（处理可变参数）
- `GenerateVarDefinitions` - 创建 "var 名称 类型" 语句
- `GroupVarsByKindToLines` - 将相同类型的变量分组
- `FormatAddressableNames` - 为名称添加 "&" 前缀
- `SimpleMakeNameFunction` - 自动命名生成（arg0、arg1、res0、res1）

**使用场景：**
- 生成具有相同签名的包裹函数
- 从接口创建模拟实现
- 提取函数签名来编写文档
- 生成带类型变量的测试设置代码

### syntaxgo_reflect - 类型信息提取

**核心函数：**
- `GetPkgPath/GetPkgPathV2` - 从类型/泛型类型获取包路径
- `GetPkgName/GetPkgNameV2` - 从类型提取包名
- `GetPkgPaths` - 批量获取多个类型的路径
- `GetTypes` - 从对象获取反射类型
- `GenerateTypeUsageCode` - 生成限定类型名（pkg.Type）
- `GetQuotedPackageImportPaths` - 创建带引号的导入路径

**使用场景：**
- 根据使用的类型自动生成导入语句
- 在代码生成中创建限定类型名
- 从类型发现包依赖
- 生成带正确导入的类型安全代码

### syntaxgo_search - AST 搜索与导航

**核心函数：**
- `FindFunctionByName` - 在 AST 中定位函数声明
- `FindStructTypeByName` - 查找结构体定义
- `FindInterfaceTypeByName` - 查找接口定义
- `FindClassesAndFunctions` - 提取顶层声明
- `GetStructFields` - 获取结构体的所有字段
- `GetStructFieldNames` - 提取字段名称
- `GetInterfaceMethods` - 列出接口方法
- `GetArrayElementType` - 获取数组/切片的元素类型

**使用场景：**
- 分析代码时查找函数
- 在 ORM 映射中提取结构体定义
- 列出接口方法来检查实现
- 导航复杂的 AST 结构

### syntaxgo_tag - 结构体标签操作

**核心函数：**
- `ExtractTagValue` - 获取完整的标签内容（如 `gorm:"column:id;type:bigint"`）
- `ExtractTagField` - 提取字段值（如 `column` → `id`）
- `ExtractTagValueIndex/ExtractTagFieldIndex` - 获取值和位置信息
- `SetTagFieldValue` - 更新/插入标签字段
- `ExtractNoValueFieldNameIndex` - 查找标志如 `primaryKey`
- `ExtractFieldEqualsValueIndex` - 查找具有特定值的字段

**使用场景：**
- 解析结构体定义中的 GORM/JSON 标签
- 生成数据库列映射
- 在代码生成中更新结构体标签
- 在 linter 中验证标签格式

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/syntaxgo.svg?variant=adaptive)](https://starchart.cc/yyle88/syntaxgo)
