package i18n

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strings"
)

type mapTranslator struct {
	c                 context.Context
	defaultLocale     string
	translations      map[string]map[string]string
	templatesByLocale map[string]*template.Template
}

func (t mapTranslator) TranslateWithMap(key, locale string, args map[string]string) string {
	s := t._translate(true, key, locale)
	return placeMapValues(s, args)
}

// NewMapTranslator creates new map translator
func NewMapTranslator(c context.Context, defaultLocale string, translations map[string]map[string]string) Translator {
	return mapTranslator{
		c:                 c,
		defaultLocale:     defaultLocale,
		translations:      translations,
		templatesByLocale: make(map[string]*template.Template),
	}
}

type theSingleLocaleTranslator struct {
	locale Locale
	Translator
}

func placeMapValues(s string, args map[string]string) string {
	for k, v := range args {
		if placeholder := "{" + k + "}"; strings.Contains(s, placeholder) {
			s = strings.ReplaceAll(s, placeholder, v)
		} else {
			s += fmt.Sprintf(" $EXTRA(%s)", k)
		}
	}
	return s
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

func (t mapTranslator) _translate(warn bool, key, locale string, args ...any) string {
	s, found := t.translations[key][locale]
	if !found {
		if warn {
			warningf(t.c, "Translation not found by key & locale: key=%v&locale=%v", key, locale)
		}
		if defaultLocale := t.defaultLocale; defaultLocale != locale {
			if defaultLocale == "" {
				defaultLocale = "en-US"
			}
			if s, found = t.translations[key][defaultLocale]; !found {
				if warn {
					warningf(t.c, "Translation not found for default locale: key=%v&locale=%v", key, defaultLocale)
				}
				return key
			}
		}
	}
	if len(args) > 0 {
		if len(args) == 1 && strings.Contains(s, "}}") && (strings.Contains(s, "{{.") || strings.Contains(s, "{{ .")) {
			tk := locale + key
			tmpl, ok := t.templatesByLocale[tk]
			if !ok {
				var err error
				if tmpl, err = template.New(key).Parse(s); err != nil {
					panic(fmt.Sprintf("Failed to parse template '%v' for locale '%v': %v", key, locale, err.Error()))
				}
				t.templatesByLocale[tk] = tmpl
			}
			var buffer bytes.Buffer
			if err := tmpl.Execute(&buffer, args[0]); err != nil {
				panic(fmt.Sprintf("Failed to render template '%v' for locale '%v': %v", key, locale, err.Error()))
			} else {
				return buffer.String()
			}
		}
		s = fmt.Sprintf(s, args...)
	}
	return s
}

func (t mapTranslator) Translate(key, locale string, args ...any) string {
	return t._translate(true, key, locale, args...)
}

func (t mapTranslator) TranslateNoWarning(key, locale string, args ...any) string {
	return t._translate(false, key, locale, args...)
}
