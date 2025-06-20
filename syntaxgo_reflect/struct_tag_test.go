package syntaxgo_reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestNewStructTag_Get(t *testing.T) {
	structTag := NewStructTag(`json:"name,omitempty"`)
	value := structTag.Get("json")
	t.Log(value)
}

func TestNewStructTag_Lookup(t *testing.T) {
	structTag := NewStructTag(`gorm:"column:product_name;type:varchar(255);comment:产品名称"`)
	value, ok := structTag.Lookup("gorm")
	require.True(t, ok)
	t.Log(value)
}
