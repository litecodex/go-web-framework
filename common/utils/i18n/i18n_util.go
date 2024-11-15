package i18n

import (
	"encoding/json"
	"github.com/litecodex/go-web-framework/web/logger"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type DatabaseI18nMessageSource interface {
	GetMessage(i18nKey string, args []interface{}, lang string) string
}

type I18nMessageTool struct {
	fileSource     *i18n.Bundle
	databaseSource DatabaseI18nMessageSource
}

var (
	EmptyArgs []interface{}
	SEPARATOR = ":"
)

func NewI18nMessageTool(filePathList []string) *I18nMessageTool {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", unmarshalFunc)
	// 读取文件内容
	for _, path := range filePathList {
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			logger.GetLogger().Warn("Error loading message file")
			logger.GetLogger().Error(err.Error())
		}
	}
	return &I18nMessageTool{fileSource: bundle}
}

func (tool *I18nMessageTool) Translate(message string, args []interface{}, lang string) string {
	if tool.fileSource == nil && tool.databaseSource == nil {
		return getDefaultValue(message)
	}

	i18nKey := getI18nKey(message)
	if i18nKey == "" {
		return message
	}

	if lang == "" {
		lang = language.English.String()
	}

	// Query from database first
	if tool.databaseSource != nil {
		i18nValue := tool.databaseSource.GetMessage(i18nKey, args, lang)
		if i18nValue != "" {
			return i18nValue
		}
	}

	// Query from i18n files
	localizer := i18n.NewLocalizer(tool.fileSource, lang)
	i18nValue, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: i18nKey,
		TemplateData: map[string]interface{}{
			"Args": args,
		},
	})

	if err == nil {
		return i18nValue
	}

	return getDefaultValue(message)
}

func (tool *I18nMessageTool) TranslateSimple(message string, language string) string {
	return tool.Translate(message, EmptyArgs, language)
}

func getI18nKey(message string) string {
	if !strings.HasPrefix(message, "${") {
		return ""
	}
	firstIndexOfSeparator := strings.Index(message, SEPARATOR)
	if firstIndexOfSeparator == -1 {
		return message[2 : len(message)-1]
	}
	return message[2:firstIndexOfSeparator]
}

func getDefaultValue(message string) string {
	if !strings.HasPrefix(message, "${") {
		return message
	}
	firstIndexOfSeparator := strings.Index(message, SEPARATOR)
	if firstIndexOfSeparator == -1 {
		return message
	}
	return message[firstIndexOfSeparator+1 : len(message)-1]
}

func unmarshalFunc(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
