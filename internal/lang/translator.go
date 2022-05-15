package lang

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var FS embed.FS

type Translator struct {
	Lang       language.Tag
	localizers map[language.Tag]*i18n.Localizer
}

func NewTranslator() *Translator {
	t := &Translator{}
	t.localizers = map[language.Tag]*i18n.Localizer{}
	return t
}

// Добавить новую локализацию в транслатор
func (t *Translator) AddLocalizer(path string, lang language.Tag) error {
	// Load translation file
	file, err := FS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("loading '%s' translation: %s", lang.String(), err)
	}

	bundle := i18n.NewBundle(lang)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = bundle.ParseMessageFileBytes(file, path)
	if err != nil {
		return err
	}

	t.localizers[lang] = i18n.NewLocalizer(bundle, lang.String())

	return nil
}

func (t Translator) TransData(str string, data interface{}) (string, error) {
	translation, err := t.localizers[t.Lang].Localize(&i18n.LocalizeConfig{
		MessageID:    str,
		TemplateData: data,
	})
	if err != nil {
		return "", fmt.Errorf("translating string '%s' with data: %s", str, err)
	}

	return translation, nil
}
