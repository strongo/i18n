package i18n

import "errors"

var _ LocalesProvider = (*supported)(nil)

type supported struct {
	locales []Locale
}

func (s supported) GetLocaleByCode5(code5 string) (Locale, error) {
	if code5 == "" {
		return LocaleUndefined, errors.New("GetLocaleByCode5(code5 string) - code5 is empty string")
	}
	for _, locale := range s.locales {
		if codeLen := len(code5); codeLen == 5 && locale.Code5 == code5 || codeLen == 2 && locale.Code5[:2] == code5 {
			return locale, nil
		}
	}
	return LocaleUndefined, errors.New("locale not found by code5: " + code5)
}

func (s supported) SupportedLocales() []Locale {
	locales := make([]Locale, len(s.locales))
	copy(locales, s.locales)
	return locales
}
