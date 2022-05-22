package lang

import (
	"golang.org/x/text/language"
	"testing"
)

func TestLang(t *testing.T) {
	d := NewTranslator()
	d.Lang = language.Russian

	err := d.AddLocalizer("langs/RU.json", language.Russian)
	if err != nil {
		t.Error(err)
	}

	text, err := d.Trans("login")
	if err != nil {
		t.Error(err)
	}
	if text == "Войти в игру" {
		t.Log("Translator test ok")
	} else {
		t.Error("translator fail")
	}

}
