package syntaxgo_reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewStructTag tests accessing struct field tags via reflection
// Verifies struct tags can be accessed and parsed from field metadata
//
// TestNewStructTag 测试通过反射访问结构体字段标签
// 验证可以从字段元数据访问和解析结构体标签
func TestNewStructTag(t *testing.T) {
	type Example struct {
		Name string `json:"name"`
		Type string `form:"type" json:"type"`
	}

	objectType := reflect.TypeOf(Example{})
	for i := 0; i < objectType.NumField(); i++ {
		field := objectType.Field(i)
		t.Log(field.Name, field.Tag.Get("json"))
	}
}

// TestNewStructTag_Get tests retrieving tag values with Get method
// Verifies Get can extract tag content from tag string
//
// TestNewStructTag_Get 测试使用 Get 方法获取标签值
// 验证 Get 能够从标签字符串提取标签内容
func TestNewStructTag_Get(t *testing.T) {
	structTag := NewStructTag(`json:"name,omitempty"`)
	value := structTag.Get("json")
	t.Log(value)
}

// TestNewStructTag_Lookup tests retrieving tag values with Lookup method
// Verifies Lookup can check tag existence and extract content
//
// TestNewStructTag_Lookup 测试使用 Lookup 方法获取标签值
// 验证 Lookup 能够检查标签存在性并提取内容
func TestNewStructTag_Lookup(t *testing.T) {
	structTag := NewStructTag(`gorm:"column:product_name;type:varchar(255);comment:产品名称"`)
	value, ok := structTag.Lookup("gorm")
	require.True(t, ok)
	t.Log(value)
}
