package strings_util

import (
	"strconv"
)

func MustToInt64(data string) int64 {
	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func MustParseBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return b
}
