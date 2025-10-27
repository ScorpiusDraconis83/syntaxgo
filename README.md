[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/syntaxgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/syntaxgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/syntaxgo)](https://pkg.go.dev/github.com/yyle88/syntaxgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/syntaxgo/master.svg)](https://coveralls.io/github/yyle88/syntaxgo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/syntaxgo.svg)](https://github.com/yyle88/syntaxgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/syntaxgo)](https://goreportcard.com/report/github.com/yyle88/syntaxgo)

# syntaxgo

**syntaxgo** is a Go toolkit built on `go/ast` Abstract Syntax Tree and `reflect` package.

**syntaxgo** is designed to ease code analysis and automated generation.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎯 **AST Parsing**: Parse Go source code into Abstract Syntax Tree with multiple parsing modes
⚡ **Type Reflection**: Extract package paths, type info, and generate usage code
🔍 **Code Search**: Find functions, types, and struct fields in AST
📝 **Tag Operations**: Extract and manipulate struct field tags
🛠️ **Code Generation**: Generate function params, return values, and variable definitions
📦 **Import Management**: Auto inject missing imports and manage package paths

## Installation

```bash
go get github.com/yyle88/syntaxgo
```

## Usage

### Basic AST Parsing

```go
package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func main() {
	// Parse Go source file from current path
	astBundle, err := syntaxgo_ast.NewAstBundleV4(runpath.Current())
	if err != nil {
		panic(err)
	}

	// Extract package name from AST
	pkgName := astBundle.GetPackageName()
	fmt.Println("Package name:", pkgName)

	// Print complete AST structure to console
	must.Done(astBundle.Print())
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Type Reflection

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type MyStruct struct {
	Name string // User name field
	Age  int    // User age field
}

func main() {
	// Extract package path from type instance
	pkgPath := syntaxgo_reflect.GetPkgPath(MyStruct{})
	fmt.Println("Package path:", pkgPath)

	// Generate qualified type usage code
	typeCode := syntaxgo_reflect.GenerateTypeUsageCode(reflect.TypeOf(MyStruct{}))
	fmt.Println("Type usage code:", typeCode)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Search in AST

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
	// Parse Go source code from string
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
	astFunc := syntaxgo_search.FindFunctionByName(astFile, "HelloWorld")
	if astFunc != nil {
		fmt.Println("Found function:", astFunc.Name.Name)
	}

	// Search struct type definition in AST
	structType, ok := syntaxgo_search.FindStructTypeByName(astFile, "Person")
	must.OK(ok)

	// Traverse struct fields and print field names
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			fmt.Println("Person Field:", fieldName.Name)
		}
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

### Tag Operations

```go
package main

import (
	"fmt"

	"github.com/yyle88/syntaxgo/syntaxgo_tag"
)

func main() {
	// Define struct field tag with GORM and JSON annotations
	tag := `gorm:"column:id;type:bigint" json:"id"`

	// Extract GORM tag value from complete tag string
	gormTag := syntaxgo_tag.ExtractTagValue(tag, "gorm")
	fmt.Println("GORM tag:", gormTag)

	// Extract column field value from GORM tag
	column := syntaxgo_tag.ExtractTagField(gormTag, "column", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column name:", column)

	// Extract type field value from GORM tag
	typeValue := syntaxgo_tag.ExtractTagField(gormTag, "type", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column type:", typeValue)
}
```

⬆️ **Source:** [Source](internal/demos/demo4x/main.go)

## Modules

### syntaxgo_ast - AST Parsing and Management

**Core Functions:**
- `NewAstBundleV1/V2/V3/V4/V5/V6` - Parse Go source files with different input modes (bytes, path, FileSet)
- `FormatSource` - Format AST back to formatted Go source code
- `SerializeAst` - Serialize AST into text representation
- `GetPackageName` - Extract package name from AST
- `AddImport/DeleteImport` - Manage single import paths
- `AddNamedImport/DeleteNamedImport` - Handle aliased imports (like `_ "embed"`)
- `InjectImports` - Auto inject missing imports into source code
- `CreateImports` - Generate import block from package paths

**Use Cases:**
- Parse source files with comment preservation
- Auto fix missing imports in generated code
- Extract package info when scanning projects
- Format code after AST manipulation

### syntaxgo_astnode - AST Node Operations

**Core Functions:**
- `NewNode/NewNodeV1/NewNodeV2` - Create nodes with position info
- `GetCode/GetText` - Extract code content from AST nodes
- `SdxEdx` - Get start/end indices of nodes
- `DeleteNodeCode` - Remove node code from source
- `ChangeNodeCode` - Replace node content with new code
- `ChangeNodeCodeSetSomeNewLines` - Replace with added newlines

**Use Cases:**
- Extract function bodies from AST
- Replace struct field definitions
- Delete unused code sections
- Insert new code at specific positions

### syntaxgo_astnorm - Function Signature Processing

**Core Functions:**
- `NewNameTypeElements` - Extract params/returns from AST field list
- `ExtractNameTypeElements` - Process field lists with custom naming
- `Names/Kinds` - Get param names and types as separate lists
- `FormatNamesWithKinds` - Format as "name type" pairs
- `GenerateFunctionParams` - Create params when calling functions (handle variadic)
- `GenerateVarDefinitions` - Create "var name type" statements
- `GroupVarsByKindToLines` - Group vars with same type
- `FormatAddressableNames` - Add "&" prefix to names
- `SimpleMakeNameFunction` - Auto name generator (arg0, arg1, res0, res1)

**Use Cases:**
- Generate wrapping functions with same signature
- Create mock implementations from interfaces
- Extract function signatures to document APIs
- Generate test setup code with typed variables

### syntaxgo_reflect - Type Information Extraction

**Core Functions:**
- `GetPkgPath/GetPkgPathV2` - Get package path from type/generic type
- `GetPkgName/GetPkgNameV2` - Extract package name from type
- `GetPkgPaths` - Batch get paths from multiple types
- `GetTypes` - Get reflect types from objects
- `GenerateTypeUsageCode` - Generate qualified type name (pkg.Type)
- `GetQuotedPackageImportPaths` - Create quoted import paths

**Use Cases:**
- Auto generate import statements based on used types
- Create qualified type names in code generation
- Discover package dependencies from types
- Generate type-safe code with proper imports

### syntaxgo_search - AST Search and Navigation

**Core Functions:**
- `FindFunctionByName` - Locate function declaration in AST
- `FindStructTypeByName` - Find struct definition
- `FindInterfaceTypeByName` - Find interface definition
- `FindClassesAndFunctions` - Extract top-most declarations
- `GetStructFields` - Get all fields from struct
- `GetStructFieldNames` - Extract field names
- `GetInterfaceMethods` - List interface methods
- `GetArrayElementType` - Get element type of arrays/slices

**Use Cases:**
- Find functions when analyzing code
- Extract struct definitions in ORM mapping
- List interface methods to check implementations
- Navigate complex AST structures

### syntaxgo_tag - Struct Tag Manipulation

**Core Functions:**
- `ExtractTagValue` - Get complete tag content (e.g., `gorm:"column:id;type:bigint"`)
- `ExtractTagField` - Extract field value (e.g., `column` → `id`)
- `ExtractTagValueIndex/ExtractTagFieldIndex` - Get values with position info
- `SetTagFieldValue` - Update/insert tag fields
- `ExtractNoValueFieldNameIndex` - Find flags like `primaryKey`
- `ExtractFieldEqualsValueIndex` - Find fields with specific values

**Use Cases:**
- Parse GORM/JSON tags in struct definitions
- Generate database column mappings
- Update struct tags in code generation
- Validate tag formats in linters

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/syntaxgo.svg?variant=adaptive)](https://starchart.cc/yyle88/syntaxgo)

---
