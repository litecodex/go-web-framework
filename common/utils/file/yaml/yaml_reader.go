package yaml

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

// ReadYaml 读取yaml格式的配置文件
func ReadYaml(path string, config interface{}) error {
	fmt.Println("Load config Path: ", path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	//fmt.Println(string(bytes)) // 打印读取到的配置信息，方便核对数据
	return yaml.Unmarshal(bytes, config)
}

func MustReadYaml(path string, config interface{}) {
	err := ReadYaml(path, config)
	if err != nil {
		panic(err)
	}
}
