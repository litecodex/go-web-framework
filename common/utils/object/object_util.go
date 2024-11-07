package object

import (
	JSON "github.com/litecodex/go-web-framework/common/utils/json"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 转为int
func ToIntValue(data interface{}) int {
	switch result := data.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		if d := ToString(data); d != "" {
			value, _ := strconv.Atoi(d)
			return value
		}
	}
	return 0
}

// 驼峰转下划线
func camelCaseToSnakeCase(s string) string {
	reg := regexp.MustCompile("([a-z])([A-Z])")
	snake := reg.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

// 转出数据库更新数据数据时使用的map
func ToDBUpdateMap(data interface{}) *map[string]interface{} {
	tmpMap := JSON.MustParseToMap(data)
	dbUpdateMap := make(map[string]interface{}, len(tmpMap))
	for key, value := range tmpMap {
		dbUpdateMap[camelCaseToSnakeCase(key)] = value
	}
	delete(dbUpdateMap, "id") // 主键不能更新
	return &dbUpdateMap
}

// 转为string
func ToString(obj interface{}) string {
	var result string
	if obj == nil {
		return result
	}

	switch obj.(type) {
	case float64:
		ft := obj.(float64)
		result = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := obj.(float32)
		result = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := obj.(int)
		result = strconv.Itoa(it)
	case uint:
		it := obj.(uint)
		result = strconv.Itoa(int(it))
	case int8:
		it := obj.(int8)
		result = strconv.Itoa(int(it))
	case uint8:
		it := obj.(uint8)
		result = strconv.Itoa(int(it))
	case int16:
		it := obj.(int16)
		result = strconv.Itoa(int(it))
	case uint16:
		it := obj.(uint16)
		result = strconv.Itoa(int(it))
	case int32:
		it := obj.(int32)
		result = strconv.Itoa(int(it))
	case uint32:
		it := obj.(uint32)
		result = strconv.Itoa(int(it))
	case int64:
		it := obj.(int64)
		result = strconv.FormatInt(it, 10)
	case uint64:
		it := obj.(uint64)
		result = strconv.FormatUint(it, 10)
	case string:
		result = obj.(string)
	case time.Time:
		t, _ := obj.(time.Time)
		result = t.String()
		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		result = strings.Replace(result, " +0800 CST", "", 1)
		result = strings.Replace(result, " +0000 UTC", "", 1)
	case []byte:
		result = string(obj.([]byte))
	default:
		result = JSON.MustStringify(obj)
	}
	return result
}
