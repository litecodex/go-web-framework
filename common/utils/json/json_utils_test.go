package json_util

import (
	"fmt"
	"testing"
)

func TestParseToMap(t *testing.T) {
	myMap := MustParseToMap("{\"key1\":\"value1\"}")
	fmt.Println(myMap)
}

func TestStringify(t *testing.T) {
	myMap := make(map[string]string)
	myMap["TestKey"] = "testValue"
	fmt.Println(MustStringify(myMap))
}
