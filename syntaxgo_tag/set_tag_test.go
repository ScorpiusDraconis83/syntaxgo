package syntaxgo_tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSetTagFieldValue_InsertFieldTop tests inserting field at top of tag
// Verifies SetTagFieldValue can insert new field at the beginning
//
// TestSetTagFieldValue_InsertFieldTop 测试在标签顶部插入字段
// 验证 SetTagFieldValue 能够在开头插入新字段
func TestSetTagFieldValue_InsertFieldTop(t *testing.T) {
	tag := `gorm:"column:id"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := INSERT_LOCATION_TOP

	expected := `gorm:"type:bigint;column:id"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

// TestSetTagFieldValue_InsertFieldEnd tests inserting field at end of tag
// Verifies SetTagFieldValue can append new field at the end
//
// TestSetTagFieldValue_InsertFieldEnd 测试在标签末尾插入字段
// 验证 SetTagFieldValue 能够在末尾追加新字段
func TestSetTagFieldValue_InsertFieldEnd(t *testing.T) {
	tag := `gorm:"column:id"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := INSERT_LOCATION_END

	expected := `gorm:"column:id;type:bigint;"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

// TestSetTagFieldValue_UpdateExistingField tests updating existing field value
// Verifies SetTagFieldValue can modify value of existing field
//
// TestSetTagFieldValue_UpdateExistingField 测试更新现有字段值
// 验证 SetTagFieldValue 能够修改现有字段的值
func TestSetTagFieldValue_UpdateExistingField(t *testing.T) {
	tag := `gorm:"column:id;type:int"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := INSERT_LOCATION_END

	expected := `gorm:"column:id;type:bigint;"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

func TestSetTagFieldValue_InsertWithNoSemicolonBeforeEnd(t *testing.T) {
	tag := `gorm:"column:id;type:int"`
	key := "gorm"
	field := "index"
	value := "idx_name"
	insertLocation := INSERT_LOCATION_END

	expected := `gorm:"column:id;type:int;index:idx_name;"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

func TestSetTagFieldValue_InvalidInsertLocation(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Log("success")
		} else {
			t.Errorf("The code did not panic on invalid insert location")
		}
	}()

	tag := `gorm:"column:id"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := InsertLocation("INVALID")

	SetTagFieldValue(tag, key, field, value, insertLocation)
}

func TestSetTagFieldValue_FieldNotFoundPanic(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Log("success")
		} else {
			t.Errorf("The code did not panic when field index was not found")
		}
	}()

	tag := `json:"id"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := INSERT_LOCATION_END

	SetTagFieldValue(tag, key, field, value, insertLocation)
}

func TestSetTagFieldValue_InsertMultipleFields(t *testing.T) {
	tag := `gorm:"column:id"`
	key := "gorm"
	field1 := "type"
	value1 := "int"
	field2 := "size"
	value2 := "10"
	insertLocation := INSERT_LOCATION_TOP

	expected := `gorm:"size:10;type:int;column:id"`

	result := SetTagFieldValue(tag, key, field1, value1, insertLocation)
	result = SetTagFieldValue(result, key, field2, value2, insertLocation)
	require.Equal(t, expected, result)
}

func TestSetTagFieldValue_InsertWithWhitespaceBeforeSemicolon(t *testing.T) {
	tag := `gorm:"column:id ;type:int"`
	key := "gorm"
	field := "index"
	value := "idx_name"
	insertLocation := INSERT_LOCATION_END

	expected := `gorm:"column:id ;type:int;index:idx_name;"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

func TestSetTagFieldValue_InsertFieldWithExistingWhitespace(t *testing.T) {
	tag := `gorm:" column:id"`
	key := "gorm"
	field := "type"
	value := "bigint"
	insertLocation := INSERT_LOCATION_TOP

	expected := `gorm:"type:bigint; column:id"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}

func TestSetTagFieldValue_InsertWithMultipleSemicolons(t *testing.T) {
	tag := `gorm:"column:id;;;type:int"`
	key := "gorm"
	field := "index"
	value := "idx_name"
	insertLocation := INSERT_LOCATION_END

	expected := `gorm:"column:id;;;type:int;index:idx_name;"`

	result := SetTagFieldValue(tag, key, field, value, insertLocation)
	require.Equal(t, expected, result)
}
