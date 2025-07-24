package i18n

// Translator provides translations by key and locale
type Translator interface {
	Translate(key, locale string, args ...any) string
	TranslateWithMap(key, locale string, args map[string]string) string
	TranslateNoWarning(key, locale string, args ...any) string
}

// SingleLocaleTranslator should be implemente by translators to a single language
type SingleLocaleTranslator interface {
	Locale() Locale
	Translate(key string, args ...any) string
	TranslateWithMap(key string, args map[string]string) string
	TranslateNoWarning(key string, args ...any) string
}

// LocalesProvider provides locale by code
type LocalesProvider interface {
	SupportedLocales() []Locale
	GetLocaleByCode5(code5 string) (Locale, error)
}
