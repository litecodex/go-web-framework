package logger

import (
	"github.com/litecodex/go-web-framework/common/utils/i18n"
)

// 全局静态变量
var i18nMessageTool *i18n.I18nMessageTool = i18n.NewI18nMessageTool([]string{})

func SetI18n(tool *i18n.I18nMessageTool) {
	i18nMessageTool = tool
}

func GetI18n() *i18n.I18nMessageTool {
	return i18nMessageTool
}
