package syntaxgo_reflect

import "reflect"

func NewStructTag(tag string) reflect.StructTag {
	return reflect.StructTag(tag)
}
