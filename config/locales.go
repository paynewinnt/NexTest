package config

/*
import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// 定义全局的本地化器
var Localizer *i18n.Localizer

func init() {
	// 创建一个新的本地化捆绑
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// 加载翻译文件
	bundle.MustLoadMessageFile("locales/en.json")
	bundle.MustLoadMessageFile("locales/es.json")
	// 添加其他语言的翻译文件...

	// 创建本地化器
	Localizer = i18n.NewLocalizer(bundle, "en")
}

// T 是一个辅助函数，用于本地化字符串
func T(messageID string, templateData map[string]string) string {
	translation, err := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		return messageID // 回退到默认消息
	}
	return translation
}
*/
