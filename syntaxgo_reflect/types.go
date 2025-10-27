package syntaxgo_reflect

import (
	"path/filepath"
	"reflect"
)

// GetType returns the reflect.Type of the provided object.
// GetType 返回提供对象的 reflect.Type。
func GetType(a any) reflect.Type {
	return reflect.TypeOf(a)
}

// GetTypeV2 is a generic version of GetType. It returns the reflect.Type of the result from GetObject[T]().
// GetTypeV2 是 GetType 的泛型版本，返回 GetObject[T]() 结果的 reflect.Type。
func GetTypeV2[T any]() reflect.Type {
	return GetType(GetObject[T]())
}

// GetTypeV3 returns the reflect.Type of an object. It checks if the object is a pointer, and if so, it returns the type of the underlying value.
// GetTypeV3 获取对象的 reflect.Type。如果对象是指针，返回指针所指向的值的类型。
func GetTypeV3(object any) reflect.Type {
	if vtp := reflect.TypeOf(object); vtp.Kind() == reflect.Ptr {
		// Elem() returns the element type for pointer types.
		// 如果对象是指针，则调用 Elem() 获取指针指向的类型。
		return vtp.Elem() // Elem() panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	} else {
		return vtp
	}
}

// GetTypeV4 is a generic function that takes a pointer and returns the reflect.Type of the dereferenced value.
// GetTypeV4 是一个泛型函数，接受一个指针并返回解引用后的值的 reflect.Type。
func GetTypeV4[T any](p *T) reflect.Type {
	return reflect.TypeOf(*p)
}

// GetTypeName returns the name of the type
// Returns the simple type name without package path
//
// GetTypeName 返回类型名称
// 返回不带包路径的简单类型名称
func GetTypeName(a any) string {
	return reflect.TypeOf(a).Name()
}

// GetTypeNameV2 is a generic version of GetTypeName
// Returns the type name using generic type param
//
// GetTypeNameV2 是 GetTypeName 的泛型版本
// 使用泛型类型参数返回类型名称
func GetTypeNameV2[T any]() string {
	return GetTypeName(GetObject[T]())
}

// GetTypeNameV3 returns type name adapting to both objects and pointers
// Handles both value types and pointer types in the same way
//
// GetTypeNameV3 返回类型名称，适配对象和指针
// 以相同方式处理值类型和指针类型
func GetTypeNameV3(object any) string {
	return GetTypeV3(object).Name()
}

// GetTypeNameV4 returns type name from a pointer using generics
// Dereferences the pointer to get the underlying type name
//
// GetTypeNameV4 使用泛型从指针返回类型名称
// 解引用指针以获取底层类型名称
func GetTypeNameV4[T any](p *T) string {
	return GetTypeV4(p).Name()
}

// GetPkgPath returns the package path of the type using reflection
// Note: Works with struct objects but not with pointers/functions/methods
// Functions cannot be reflected to get package path - use direct code reference instead
//
// GetPkgPath 使用反射返回类型的包路径
// 注意：适用于结构体对象但不适用于指针/函数/方法类型
// 函数无法通过反射获取包路径 - 请使用直接代码引用
func GetPkgPath(a any) string {
	return reflect.TypeOf(a).PkgPath()
}

// GetPkgPathV2 is a generic version of GetPkgPath
// Both versions are useful: non-generic when passed from non-generic callers
// Generic version adds type safety using just V2 suffix
//
// GetPkgPathV2 是 GetPkgPath 的泛型版本
// 两个版本都有用：非泛型版用于非泛型调用者传递对象时
// 泛型版本通过 V2 后缀增加类型安全性
func GetPkgPathV2[T any]() string {
	return reflect.TypeOf(GetObject[T]()).PkgPath()
}

// GetPkgPathV3 returns package path adapting to both objects and pointers
// Handles both value types and pointer types in the same way
//
// GetPkgPathV3 返回包路径，适配对象和指针
// 以相同方式处理值类型和指针类型
func GetPkgPathV3(a any) string {
	return GetTypeV3(a).PkgPath()
}

// GetPkgPathV4 returns package path from a pointer using generics
// Dereferences the pointer to get the underlying type's package path
//
// GetPkgPathV4 使用泛型从指针返回包路径
// 解引用指针以获取底层类型的包路径
func GetPkgPathV4[T any](p *T) string {
	return GetTypeV4(p).PkgPath()
}

// GetPkgName extracts the package name from the package path
// Returns the base name of the package path (last segment)
//
// GetPkgName 从包路径提取包名
// 返回包路径的基础名称（最后一段）
func GetPkgName(a any) string {
	var pkgPath = GetPkgPath(a)
	if pkgPath == "" {
		return ""
	}
	return filepath.Base(pkgPath)
}

// GetPkgNameV2 is a generic version of GetPkgName
// Returns the package name using generic type param
//
// GetPkgNameV2 是 GetPkgName 的泛型版本
// 使用泛型类型参数返回包名
func GetPkgNameV2[T any]() string {
	return GetPkgName(GetObject[T]())
}

// GetPkgNameV3 extracts package name adapting to both objects and pointers
// Handles both value types and pointer types in the same way
//
// GetPkgNameV3 提取包名，适配对象和指针
// 以相同方式处理值类型和指针类型
func GetPkgNameV3(a any) string {
	var pkgPath = GetPkgPathV3(a)
	if pkgPath == "" {
		return ""
	}
	return filepath.Base(pkgPath)
}

// GetPkgNameV4 extracts package name from a pointer using generics
// Dereferences the pointer to get the underlying type's package name
//
// GetPkgNameV4 使用泛型从指针提取包名
// 解引用指针以获取底层类型的包名
func GetPkgNameV4[T any](p *T) string {
	var pkgPath = GetPkgPathV4(p)
	if pkgPath == "" {
		return ""
	}
	return filepath.Base(pkgPath)
}
