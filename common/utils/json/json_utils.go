package json_util

import (
	"github.com/goccy/go-json"
)

// 序列化
func MustStringify(data interface{}) string {
	result, err := Stringify(data)
	if err != nil {
		panic(err)
	}
	return result
}

func Stringify(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// 反序列化
func MustParseToMap(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	if jsonString, ok := obj.(string); ok {
		// 使用 json.Unmarshal 将 JSON 字符串解析到 map 中
		err := json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			panic(err)
		}
		return result
	} else {
		MustParse(MustStringify(obj), &result)
		return result
	}
}

func MustParse(jsonString string, result interface{}) {
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		panic(err)
	}
}

func Parse(jsonString string, result interface{}) error {
	return json.Unmarshal([]byte(jsonString), &result)
}
