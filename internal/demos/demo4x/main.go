// Demo4x shows struct tag parsing and field extraction
// Parse GORM and JSON tags from struct field tags
// Extract specific fields and values from tag strings
//
// Demo4x 展示结构体标签解析和字段提取
// 从结构体字段标签解析 GORM 和 JSON 标签
// 从标签字符串提取特定字段和值
package main

import (
	"fmt"

	"github.com/yyle88/syntaxgo/syntaxgo_tag"
)

func main() {
	// Define struct field tag with GORM and JSON annotations
	// Tag format follows Go struct tag conventions
	//
	// 定义带有 GORM 和 JSON 注解的结构体字段标签
	// 标签格式遵循 Go 结构体标签约定
	tag := `gorm:"column:id;type:bigint" json:"id"`

	// Extract GORM tag value from complete tag string
	// Get the content inside gorm:"..." quotes
	//
	// 从完整标签字符串提取 GORM 标签值
	// 获取 gorm:"..." 引号内的内容
	gormTag := syntaxgo_tag.ExtractTagValue(tag, "gorm")
	fmt.Println("GORM tag:", gormTag)

	// Extract column field value from GORM tag
	// Parse column:id to get database column name
	//
	// 从 GORM 标签提取 column 字段值
	// 解析 column:id 获取数据库列名
	column := syntaxgo_tag.ExtractTagField(gormTag, "column", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column name:", column)

	// Extract type field value from GORM tag
	// Parse type:bigint to get database column type
	//
	// 从 GORM 标签提取 type 字段值
	// 解析 type:bigint 获取数据库列类型
	typeValue := syntaxgo_tag.ExtractTagField(gormTag, "type", syntaxgo_tag.EXCLUDE_WHITESPACE_PREFIX)
	fmt.Println("Column type:", typeValue)
}
