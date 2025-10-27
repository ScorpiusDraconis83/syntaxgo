package syntaxgo_reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenerateTypeUsageCode tests generating usage code for types from other packages
// Verifies GenerateTypeUsageCode produces correct qualified type names
//
// TestGenerateTypeUsageCode 测试生成从其他包使用类型的代码
// 验证 GenerateTypeUsageCode 生成正确的限定类型名
func TestGenerateTypeUsageCode(t *testing.T) {
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(reflect.TypeOf(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetType(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetTypeV2[Example]()))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetTypeV3(&Example{})))
}
