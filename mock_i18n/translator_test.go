package mock_i18n

import (
	"testing"

	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
)

var _ i18n.Translator = (*MockTranslator)(nil)

func TestTranslator(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockTranslator(ctrl)
	if m == nil {
		t.Fatal("translator is nil")
	}

	m.EXPECT().Translate("k1", "en", "arg1").Return("res1")
	if got := m.Translate("k1", "en", "arg1"); got != "res1" {
		t.Errorf("got %q, want res1", got)
	}

	m.EXPECT().TranslateNoWarning("k2", "en", "arg2").Return("res2")
	if got := m.TranslateNoWarning("k2", "en", "arg2"); got != "res2" {
		t.Errorf("got %q, want res2", got)
	}

	m.EXPECT().TranslateWithMap("k3", "en", map[string]string{"a": "b"}).Return("res3")
	if got := m.TranslateWithMap("k3", "en", map[string]string{"a": "b"}); got != "res3" {
		t.Errorf("got %q, want res3", got)
	}
}
