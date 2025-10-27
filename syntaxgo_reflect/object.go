package syntaxgo_reflect

// GetObject returns a zero value of type T
// Compile-time prevents T from being a pointer type (like *A)
//
// GetObject 返回类型 T 的零值
// 在编译阶段预防 T 是指针类型（比如 *A）的情况
func GetObject[T any]() (a T) {
	return a
}

// NewObject creates a new pointer to type T
// Returns a pointer to a fresh allocated zero value of type T
//
// NewObject 创建类型 T 的新指针
// 返回指向新分配的类型 T 零值的指针
func NewObject[T any]() (a *T) {
	return new(T)
}
