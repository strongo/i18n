package i18n

import (
	"context"
	"testing"
)

// A key with no translation must come back as the key itself for every
// locale, including the translator's own default locale. Before this test the
// default locale returned "" while every other locale returned the key, which
// is how a bot description sent through Translate came out empty.
func TestMapTranslator_MissingKeyReturnsKeyForDefaultLocale(t *testing.T) {
	tr := NewMapTranslator(context.Background(), "en-UK", map[string]map[string]string{
		"known": {"en-UK": "Known", "ru-RU": "Известно"},
	})
	const raw = "Family & friends assistant."
	for _, locale := range []string{"en-UK", "ru-RU", "de-DE"} {
		if got := tr.Translate(raw, locale); got != raw {
			t.Errorf("Translate(missing, %q) = %q, want the key back", locale, got)
		}
		if got := tr.TranslateNoWarning(raw, locale); got != raw {
			t.Errorf("TranslateNoWarning(missing, %q) = %q, want the key back", locale, got)
		}
	}
	if got := tr.Translate("known", "en-UK"); got != "Known" {
		t.Errorf("Translate(known, en-UK) = %q", got)
	}
	if got := tr.Translate("known", "ru-RU"); got != "Известно" {
		t.Errorf("Translate(known, ru-RU) = %q", got)
	}
}
