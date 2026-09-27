package mock_i18n

import (
	"testing"

	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
)

var _ i18n.SingleLocaleTranslator = (*MockSingleLocaleTranslator)(nil)

func TestNewMockSingleLocaleTranslator(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockSingleLocaleTranslator(ctrl)
	if m == nil {
		t.Fatal("translator is nil")
	}

	dummyLocale := i18n.Locale{Code5: "en-US"}
	m.EXPECT().Locale().Return(dummyLocale)
	if loc := m.Locale(); loc.Code5 != "en-US" {
		t.Fatalf("unexpected Locale: %v", loc)
	}

	m.EXPECT().Translate("k1", "arg1").Return("res1")
	if got := m.Translate("k1", "arg1"); got != "res1" {
		t.Errorf("got %q, want res1", got)
	}

	m.EXPECT().TranslateNoWarning("k2", "arg2").Return("res2")
	if got := m.TranslateNoWarning("k2", "arg2"); got != "res2" {
		t.Errorf("got %q, want res2", got)
	}

	m.EXPECT().TranslateWithMap("k3", map[string]string{"a": "b"}).Return("res3")
	if got := m.TranslateWithMap("k3", map[string]string{"a": "b"}); got != "res3" {
		t.Errorf("got %q, want res3", got)
	}
}
