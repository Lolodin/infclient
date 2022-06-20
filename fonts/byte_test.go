package fonts

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	f, err := os.Open("alundratext.ttf")
	if err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	f, err = os.Create("alundra.go")
	if err != nil {
		t.Error(err)
	}
	fmt.Fprintf(f, "var %s = []byte(%q)\n", "fonts", string(b))
}
