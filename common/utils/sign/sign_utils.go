package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
)

// MD5签名
func MD5Sign(data interface{}, signSalt string) string {
	// 使用反射获取结构体的类型和值
	valueType := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	signMap := make(map[string]interface{})
	if valueType.Kind() == reflect.Struct {
		// 遍历结构体的字段
		for i := 0; i < valueType.NumField(); i++ {
			// 获取字段名和字段值
			fieldName := valueType.Field(i).Name
			fieldValue := value.Field(i).Interface()
			// 将字段名和字段值插入到map中
			signMap[fieldName] = fieldValue
		}
	} else if valueType.Kind() == reflect.Map {
		// If data is already a map[string]interface{}, simply return it
		if valueType.ConvertibleTo(reflect.TypeOf(map[string]interface{}{})) {
			signMap = data.(map[string]interface{})
		}
		// Convert each key-value pair to map[string]interface{}
		keys := value.MapKeys()
		for _, key := range keys {
			signMap[key.String()] = value.MapIndex(key).Interface()
		}
	} else {
		return ""
	}

	// 对字段按照 ASCII 码进行排序
	delete(signMap, "MD5Sign") // sign字段不参与签名
	var signFileNames []string
	for key := range signMap {
		signFileNames = append(signFileNames, key)
	}
	sort.Strings(signFileNames)

	// 将排序后的字段拼接成一个字符串
	// 初始化一个空的切片，用于存储拼接后的键值对字符串
	signContent := ""
	// 遍历map，拼接键值对字符串，并将其添加到切片中
	for _, fileName := range signFileNames {
		// 将所有value拼接起来
		signContent += fmt.Sprintf("%v", signMap[fileName])
	}
	// 使用&符号将切片中的所有字符串连接起来
	signContent += signSalt

	// 计算MD5签名
	hasher := md5.New()
	hasher.Write([]byte(signContent))
	md5Signature := hex.EncodeToString(hasher.Sum(nil))

	return md5Signature
}
