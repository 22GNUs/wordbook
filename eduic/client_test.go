package eduic

import (
	"strings"
	"testing"

	"github.com/22GNUs/wordbook/cfg"
)

var authProvider = func() string {
	c := cfg.Read()
	return c.Auth
}

// TestClientAddWords calls api.AddWord with multiple words, check if resposne message is right
func TestClientAddWords(t *testing.T) {
	want := "单词导入成功,导入数量"

	client := NewClient(authProvider)
	msg, err := client.AddWords("test")
	if err != nil {
		t.Fatalf("AddWord(%s) error eccured, want: %s, error: %s", "test", want, err)
	}
	if !strings.HasPrefix(msg, want) {
		t.Fatalf("AddWord(%s) return %s but want: %s", "test", msg, want)
	}
}

func TestClientListWords(t *testing.T) {
	client := NewClient(authProvider)
	explains, err := client.ListWords(0, 1)
	if err != nil {
		t.Fatalf(`ListWords(0, 1) error eccured, %s`, err)
	}
	if len(explains) == 0 {
		t.Fatalf("ListWords(0, 1) should have more then one element")
	}
}
