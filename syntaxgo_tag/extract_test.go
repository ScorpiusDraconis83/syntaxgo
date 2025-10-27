package syntaxgo_tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestExtractTagValue tests extracting tag value by key
// Verifies ExtractTagValue can extract complete tag content for a given key
//
// TestExtractTagValue 测试按键提取标签值
// 验证 ExtractTagValue 能够为给定键提取完整的标签内容
func TestExtractTagValue(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)
}

// TestExtractTagField tests extracting specific field values from tag content
// Verifies ExtractTagField handles various formats and whitespace scenarios
//
// TestExtractTagField 测试从标签内容提取特定字段值
// 验证 ExtractTagField 能够处理各种格式和空白场景
func TestExtractTagField(t *testing.T) {
	t.Run("case-1", func(t *testing.T) {
		const tmp = "column:name; primaryKey;"

		field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, "name", field)
	})

	t.Run("case-2", func(t *testing.T) {
		const tmp = "column: name; primaryKey;"

		field := ExtractTagField(tmp, "column", INCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, " name", field)
	})

	t.Run("case-3", func(t *testing.T) {
		const tmp = "column: ; primaryKey;"

		field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, "", field)
	})

	t.Run("case-4", func(t *testing.T) {
		const tmp = "column:; primaryKey;"

		field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, "", field)
	})

	t.Run("case-5", func(t *testing.T) {
		const tmp = "column: "

		field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, "", field)
	})

	t.Run("case-6", func(t *testing.T) {
		const tmp = "column:"

		field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field)
		require.Equal(t, "", field)
	})
}

// TestExtract tests combined tag value and field extraction
// Verifies both ExtractTagValue and ExtractTagField work together
//
// TestExtract 测试组合标签值和字段提取
// 验证 ExtractTagValue 和 ExtractTagField 能够协同工作
func TestExtract(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)

	field := ExtractTagField(value, "column", INCLUDE_WHITESPACE_PREFIX)
	t.Log(field)
	require.Equal(t, "name", field)
}

// TestExtractTagValueIndex tests extracting tag value with position indices
// Verifies ExtractTagValueIndex returns correct value and its position in tag
//
// TestExtractTagValueIndex 测试提取标签值及其位置索引
// 验证 ExtractTagValueIndex 返回正确的值和其在标签中的位置
func TestExtractTagValueIndex(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value, sdx, edx := ExtractTagValueIndex(tag, "gorm")
	t.Log(value, sdx, edx)
	require.Equal(t, "column:name; primaryKey;", value)
	sub := tag[sdx:edx]
	require.Equal(t, value, sub)
}

// TestExtractTagFieldIndex tests extracting field value with position indices
// Verifies ExtractTagFieldIndex returns correct field and its position
//
// TestExtractTagFieldIndex 测试提取字段值及其位置索引
// 验证 ExtractTagFieldIndex 返回正确的字段和其位置
func TestExtractTagFieldIndex(t *testing.T) {
	t.Run("case-1", func(t *testing.T) {
		const tmp = "column:name; primaryKey;"

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "name", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})

	t.Run("case-2", func(t *testing.T) {
		const tmp = "column:name"

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "name", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})

	t.Run("case-3", func(t *testing.T) {
		const tmp = `column:; primaryKey;`

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})

	t.Run("case-4", func(t *testing.T) {
		const tmp = `column: ; primaryKey;`

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})

	t.Run("case-5", func(t *testing.T) {
		const tmp = `column: `

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})

	t.Run("case-6", func(t *testing.T) {
		const tmp = `column:`

		field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
		t.Log(field, sdx, edx)
		require.Equal(t, "", field)
		sub := tmp[sdx:edx]
		require.Equal(t, field, sub)
	})
}

func TestExtractTagFieldIndex_2(t *testing.T) {
	const tmp = "column: name; primaryKey;"

	field, sdx, edx := ExtractTagFieldIndex(tmp, "column", INCLUDE_WHITESPACE_PREFIX)
	t.Log(field, sdx, edx)
	require.Equal(t, " name", field)
	sub := tmp[sdx:edx]
	require.Equal(t, field, sub)
}

func TestExtractNoValueFieldNameIndex(t *testing.T) {
	const tmp = "column:name;index;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_2(t *testing.T) {
	const tmp = "column:name;index:abc;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_3(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_4(t *testing.T) {
	const tmp = "column:name;index:value;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_5(t *testing.T) {
	const tmp = "column:name; index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, " index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_6(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_7(t *testing.T) {
	const tmp = "column:name;index;field2;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_8(t *testing.T) {
	const tmp = "column:name;index;field2:value;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_9(t *testing.T) {
	const tmp = "column:name;index;field2: value;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_10(t *testing.T) {
	const tmp = "column:name;index :value;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_11(t *testing.T) {
	const tmp = "column:name;index;field2 ;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2 ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_12(t *testing.T) {
	const tmp = "column:name;index ;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_13(t *testing.T) {
	const tmp = "column:name;  index  ;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "  index  ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_14(t *testing.T) {
	const tmp = "column:name;index;field2:value"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_15(t *testing.T) {
	const tmp = "column:name;index;field2 "

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2 ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_16(t *testing.T) {
	const tmp = "column:name;index:" // 有冒号但无值，仍应视为无值，通过
	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index:", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_17(t *testing.T) {
	const tmp = "column:name;index:value" // 没有分号结尾，且有值，不通过
	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_18(t *testing.T) {
	const tmp = "column:name;  index\t :  \t ;field2" // 含制表符、空格，但无值，通过
	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "  index\t :  \t ", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex(t *testing.T) {
	const tmp = "column:name;index:idx_myname;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, "idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_2(t *testing.T) {
	const tmp = "column:name;index:idx_myname ;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, "idx_myname ", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_3(t *testing.T) {
	const tmp = "column:name;index: idx_myname;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_4(t *testing.T) {
	const tmp = "column:name;index: idx_myname;index: idx_myname2;"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname2")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname2", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_5(t *testing.T) {
	const tmp = "column:name;index: idx_myname2;index: idx_myname;"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndexV2(t *testing.T) {
	const tmp = "column:name;index: idx_myname;index: idx_myname2,priority:2;"
	//不仅以分号或结尾为分割，还以逗号为分隔符
	sdx, edx := ExtractFieldEqualsValueIndexV2(tmp, "index", "idx_myname2", []string{","})
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname2", tmp[sdx:edx])
}
