package dbtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JsonMap 是自定义的类型，用于表示 JSON 字段
type JsonMap map[string]interface{}

// Scan 实现了 sql.Scanner 接口，用于从数据库字段扫描数据到 JsonMap
func (j *JsonMap) Scan(value interface{}) error {
	// 如果数据库返回的是 NULL，直接赋值为空 map
	if value == nil {
		*j = JsonMap{}
		return nil
	}
	// 将数据库中的 JSON 字符串（[]byte）解码为 JsonMap
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		return fmt.Errorf("unsupported scan type: %T", v)
	}
}

// Value 实现了 driver.Valuer 接口，用于将 JsonMap 类型转换为数据库字段值（JSON 格式的 []byte）
func (j JsonMap) Value() (driver.Value, error) {
	// 将 JsonMap 编码为 JSON 格式的字节数组
	return json.Marshal(j)
}
