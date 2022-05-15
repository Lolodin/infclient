package lang

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLang(t *testing.T) {
	d := NewTranslator()

	err := d.AddLocalizer("langs/RU.json", language.Russian)
	if err != nil {
		t.Error(err)
	}
	text, err := d.TransData("Login", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(text)

}
