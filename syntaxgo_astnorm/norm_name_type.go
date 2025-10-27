// Package syntaxgo_astnorm provides function signature processing utilities
// Extract and process function params and return values from AST
// Generate code for function calls, variable definitions, and type conversions
//
// syntaxgo_astnorm 包提供函数签名处理工具
// 从 AST 提取和处理函数参数和返回值
// 生成函数调用、变量定义和类型转换的代码
package syntaxgo_astnorm

import (
	"go/ast"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

// NameTypeElement represents a single element with a name, type, and associated information.
// NameTypeElement 代表一个包含名称、类型和相关信息的元素
type NameTypeElement struct {
	Name       string   // Field name / 字段名称
	Kind       string   // Type description (e.g., int, string, A, utils.A, *B, *pkg.B) / 类型描述 (例如: int, string, A, utils.A, *B, *pkg.B)
	Type       ast.Expr // Go's type representation / Go 的类型表示
	IsEllipsis bool     // Indicates if the type is a variadic (e.g., ...int, ...string) / 是否是变参类型 (例如: ...int, ...string)
}

// NewNameTypeElement creates a new NameTypeElement with parsed source code info and normalized data
// Parses param/return information and prepares it to use in code generation
//
// NewNameTypeElement 创建新的 NameTypeElement，包含已解析的源代码信息和规范化数据
// 解析参数/返回信息并准备用于代码生成
func NewNameTypeElement(
	field *ast.Field, // AST field with param/return info / AST 字段包含参数/返回值信息
	paramName string, // Name of the param / 参数名称
	paramType string, // Type of the param / 参数类型
	isEllipsis bool, // Indicates if param is a variadic / 是否为变参
	packageName string, // Package name to use with outside types / 外部类型使用的包名
	genericTypeParams map[string]ast.Expr, // Map with generic type params / 泛型类型参数的映射
) *NameTypeElement {
	elem := &NameTypeElement{
		Name:       paramName,
		Kind:       paramType,
		Type:       field.Type,
		IsEllipsis: isEllipsis,
	}
	if packageName != "" {
		elem.AdjustTypeWithPackage(packageName, genericTypeParams)
	}
	return elem
}

// AdjustTypeWithPackage adjusts the type to include the package name if needed (when using outside types).
// AdjustTypeWithPackage 如果需要（对于外部类型），将类型调整为包含包名。
func (element *NameTypeElement) AdjustTypeWithPackage(
	packageName string, // The package name / 包名
	genericTypeParams map[string]ast.Expr, // Map of generic type parameters / 泛型类型参数的映射
) {
	shortKind, adjusted := adjustKindWithPackage(strings.TrimSpace(element.Kind), packageName, genericTypeParams, element.IsEllipsis)
	if !adjusted {
		return
	}
	element.Kind = shortKind
}

func adjustKindWithPackage(shortKind string, packageName string, genericTypeParams map[string]ast.Expr, isVariadic bool) (string, bool) {
	if isVariadic {
		must.True(strings.HasPrefix(shortKind, "..."))
		shortKind = strings.TrimPrefix(shortKind, "...")
		shortKind = strings.TrimSpace(shortKind)
	}

	if strings.Contains(shortKind, ".") {
		return "", false // contains package name / 已经包含包名
	}

	if s := shortKind[0]; string(s) == "*" {
		if vcc := shortKind[1]; vcc >= 'A' && vcc <= 'Z' {
			className := shortKind[1:]
			if _, ok := genericTypeParams[className]; ok {
				return "", false // It's a generic type / 是泛型类型
			}
			shortKind = "*" + packageName + "." + className
		} else {
			return "", false // basic-type(int string float64) || // not-exportable-type
		}
	} else {
		if s := shortKind[0]; s >= 'A' && s <= 'Z' {
			className := shortKind
			if _, ok := genericTypeParams[className]; ok {
				return "", false // It's a generic type / 是泛型类型
			}
			shortKind = packageName + "." + className
		} else {
			return "", false // basic-type(int string float64) || // not-exportable-type
		}
	}

	if isVariadic {
		shortKind = "..." + shortKind
	}
	return shortKind, true
}

// MakeNameFunction generates a name when processing function params and return values.
// MakeNameFunction 用于为参数或返回值生成名称。
type MakeNameFunction func(name *ast.Ident, kind string, idx int, anonymousIdx int) string

// NameTypeElements is a collection of NameTypeElement.
// NameTypeElements 是 NameTypeElement 的集合。
type NameTypeElements []*NameTypeElement

// NewNameTypeElements creates a list of NameTypeElements based on AST field list.
// NewNameTypeElements 根据 AST 字段列表创建 NameTypeElements 的列表。
func NewNameTypeElements(
	fieldList *ast.FieldList, // AST field list with function params / AST 字段列表，表示函数参数
	nameFunc MakeNameFunction, // Function to generate param names / 用于生成参数名称的函数
	source []byte, // Source code to extract information / 用于提取信息的源代码
	packageName string, // Package name when using outside types / 外部类型的包名
	genericTypeParams map[string]ast.Expr, // Map of generic type params / 泛型类型参数的映射
) NameTypeElements {
	if fieldList == nil {
		return make(NameTypeElements, 0) // No fields, new list / 返回一个空列表
	}
	return ExtractNameTypeElements(fieldList.List, nameFunc, source, packageName, genericTypeParams)
}

// ExtractNameTypeElements extracts NameTypeElements from the AST fields.
// ExtractNameTypeElements 从 AST 字段中提取 NameTypeElements。
func ExtractNameTypeElements(
	fields []*ast.Field, // List of AST fields / AST 字段列表
	nameFunc MakeNameFunction, // Function to generate names / 用于生成名称的函数
	source []byte, // Source code / 源代码
	packageName string, // Package name / 包名
	genericTypeParams map[string]ast.Expr, // Map of generic type params / 泛型类型参数的映射
) NameTypeElements {
	var elements = make(NameTypeElements, 0) // New elements list / 创建一个空的元素列表
	var anonymousCount = 0                   // Count of anonymous fields / 匿名字段计数器
	for _, field := range fields {
		var stringType string
		var isVariadic bool
		if ellipsis, ok := field.Type.(*ast.Ellipsis); ok {
			stringType = string(source[ellipsis.Ellipsis-1 : field.Type.End()-1]) // Extract ellipsis type / 提取变参类型
			isVariadic = true
		} else {
			stringType = strings.TrimSpace(string(syntaxgo_astnode.GetCode(source, field.Type))) // Extract normal type / 提取常规类型
		}
		if len(field.Names) > 0 { // Params have names, but returns often don't / 参数有名称，但返回值通常没有
			for _, fieldName := range field.Names {
				count := len(elements)
				paramName := nameFunc(fieldName, stringType, count, 0) // Generate name when processing param / 为参数生成名称
				elem := NewNameTypeElement(field, paramName, stringType, isVariadic, packageName, genericTypeParams)
				elements = append(elements, elem)
			}
		} else {
			count := len(elements)
			paramName := nameFunc(nil, stringType, count, anonymousCount) // Generate name for anonymous field / 为匿名字段生成名称
			elem := NewNameTypeElement(field, paramName, stringType, isVariadic, packageName, genericTypeParams)
			elements = append(elements, elem)
			anonymousCount++
		}
	}
	return elements
}

// Names returns a list of names of the elements.
// Names 返回元素名称的列表。
func (elements NameTypeElements) Names() StatementParts {
	var names = make([]string, 0, len(elements))
	for _, element := range elements {
		names = append(names, element.Name)
	}
	return names
}

// Kinds returns a list of types of the elements.
// Kinds 返回元素类型的列表。
func (elements NameTypeElements) Kinds() []string {
	var kinds = make([]string, 0, len(elements))
	for _, node := range elements {
		kinds = append(kinds, node.Kind)
	}
	return kinds
}

// FormatAddressableNames returns the names prefixed with "&" (addressable names).
// FormatAddressableNames 返回带有 "&" 前缀的名称（可寻址名称）。
func (elements NameTypeElements) FormatAddressableNames() StatementParts {
	return utils.SetPrefix2Strings("&", elements.Names()) // Add '&' to each name for addressable types / 为每个名称添加 "&" 使其可寻址
}

// GenerateFunctionParams generates the function parameters list.
// GenerateFunctionParams 生成函数参数列表。
func (elements NameTypeElements) GenerateFunctionParams() StatementParts {
	var params = make([]string, 0, len(elements))
	for _, element := range elements {
		if element.IsEllipsis {
			params = append(params, element.Name+"...")
		} else {
			params = append(params, element.Name)
		}
	}
	return params
}

// FormatNamesWithKinds returns the names with associated types (e.g., "a int").
// FormatNamesWithKinds 返回名称及其类型（例如："a int"）。
func (elements NameTypeElements) FormatNamesWithKinds() StatementParts {
	var results = make([]string, 0, len(elements)) // Initialize a slice for the results / 初始化一个切片来存放结果
	for _, element := range elements {
		results = append(results, element.Name+" "+element.Kind) // Combine name and type / 合并名称和类型
	}
	return results
}

// GenerateVarDefinitions generates variable definitions (e.g., "var a int").
// GenerateVarDefinitions 生成变量定义（例如："var a int"）。
func (elements NameTypeElements) GenerateVarDefinitions() StatementLines {
	return utils.SetPrefix2Strings("var ", elements.FormatNamesWithKinds()) // Add "var" to each definition / 为每个定义添加 "var"
}

// GroupVarsByKindToLines groups variables based on type and generates the definition lines.
// GroupVarsByKindToLines 按类型分组变量，并生成相应的定义行。
func (elements NameTypeElements) GroupVarsByKindToLines() StatementLines {
	var typeToNamesMap = map[string][]string{} // Mapping of type to variable names / 类型到名称的映射
	var uniqueTypes []string                   // Store unique types / 存储唯一的类型

	// Process elements and group based on type / 遍历所有元素，按类型分组
	for _, element := range elements {
		if names, exists := typeToNamesMap[element.Kind]; exists {
			typeToNamesMap[element.Kind] = append(names, element.Name)
		} else {
			typeToNamesMap[element.Kind] = []string{element.Name}
			uniqueTypes = append(uniqueTypes, element.Kind) // Record type as encountered / 按照出现顺序记录类型
		}
	}

	// Generate variable definition lines / 生成变量定义行
	var definitionLines = make([]string, 0, len(uniqueTypes))
	for _, kind := range uniqueTypes {
		names := typeToNamesMap[kind]
		parts := strings.Join(names, ", ") // Join names with commas / 用逗号连接名称
		definitionLines = append(definitionLines, "var "+parts+" "+kind)
	}

	return definitionLines
}
