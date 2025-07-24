package i18n

type theSingleLocaleTranslator struct {
	locale Locale
	Translator
}

func (t theSingleLocaleTranslator) TranslateWithMap(key string, args map[string]string) string {
	var s = t.Translator.Translate(key, t.locale.Code5)
	return placeMapValues(s, args)
}

func (t theSingleLocaleTranslator) Translate(key string, args ...any) string {
	return t.Translator.Translate(key, t.locale.Code5, args...)
}

func (t theSingleLocaleTranslator) Locale() Locale {
	return t.locale
}

func (t theSingleLocaleTranslator) TranslateNoWarning(key string, args ...any) string {
	return t.Translator.TranslateNoWarning(key, t.locale.Code5, args...)
}

var _ SingleLocaleTranslator = (*theSingleLocaleTranslator)(nil)

// NewSingleMapTranslator creates new single map translator
func NewSingleMapTranslator(locale Locale, translator Translator) SingleLocaleTranslator {
	return theSingleLocaleTranslator{
		locale:     locale,
		Translator: translator,
	}
}
