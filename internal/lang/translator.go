package lang

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"io"
	"os"
)

var FS embed.FS

type Translator struct {
	Lang       language.Tag
	localizers map[language.Tag]words
}

type words map[string]string

func NewTranslator() *Translator {
	t := &Translator{}
	t.localizers = map[language.Tag]words{}
	return t
}

// Добавить новую локализацию в транслатор
func (t *Translator) AddLocalizer(path string, lang language.Tag) error {
	// Load translation file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("loading '%s' translation: %s", lang.String(), err)
	}
	buff, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("loading '%s' translation: %s", lang.String(), err)
	}
	w := words{}
	json.Unmarshal(buff, &w)
	t.localizers[lang] = w

	return nil
}

func (t Translator) Trans(str string) (string, error) {
	translation, ok := t.localizers[t.Lang]
	if !ok {
		return "", fmt.Errorf("translating string '%s'", str)
	}

	if word, ok := translation[str]; ok {
		return word, nil
	}

	return "", errors.New("word is not found")
}

func (t Translator) TransWithOut(str string) string {
	translation, ok := t.localizers[t.Lang]
	if !ok {
		return ""
	}

	if word, ok := translation[str]; ok {
		return word
	}

	return ""
}
