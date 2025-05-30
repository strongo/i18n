package i18n

import (
	"fmt"
	"strings"
)

type TranslatorProvider = func(locale string) Translator

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

var _ SingleLocaleTranslator = (*SingleLocaleTranslatorWithBackup)(nil)

// SingleLocaleTranslatorWithBackup should be implemente by translators to a single language with backup to another one.
type SingleLocaleTranslatorWithBackup struct {
	PrimaryTranslator SingleLocaleTranslator
	BackupTranslator  SingleLocaleTranslator
}

func (t SingleLocaleTranslatorWithBackup) TranslateWithMap(key string, args map[string]string) string {
	s := t.PrimaryTranslator.TranslateWithMap(key, args)
	if s == "" {
		s = t.BackupTranslator.TranslateWithMap(key, args)
	}
	return s
}

// NewSingleLocaleTranslatorWithBackup creates SingleLocaleTranslatorWithBackup
func NewSingleLocaleTranslatorWithBackup(primary, backup SingleLocaleTranslator) SingleLocaleTranslatorWithBackup {
	return SingleLocaleTranslatorWithBackup{PrimaryTranslator: primary, BackupTranslator: backup}
}

// Locale returns local of the translator
func (t SingleLocaleTranslatorWithBackup) Locale() Locale {
	return t.PrimaryTranslator.Locale()
}

// Translate translates
func (t SingleLocaleTranslatorWithBackup) Translate(key string, args ...any) string {
	result := t.PrimaryTranslator.Translate(key, args...)
	if result == key || result == "" {
		result = t.BackupTranslator.Translate(key, args...)
	}
	if result == "" {
		result = key + fmt.Sprintf("(args=%+v)", args)
	}
	return result
}

// TranslateNoWarning translates and does not log warning if translation not found
func (t SingleLocaleTranslatorWithBackup) TranslateNoWarning(key string, args ...any) string {
	result := t.PrimaryTranslator.TranslateNoWarning(key, args...)
	if result == key {
		result = t.BackupTranslator.TranslateNoWarning(key, args...)
	}
	return result
}

// LocalesProvider provides locale by code
type LocalesProvider interface {
	SupportedLocales() []Locale
	GetLocaleByCode5(code5 string) (Locale, error)
}

// Locale describes language
type Locale struct {
	Code5        string
	NativeTitle  string
	EnglishTitle string
	FlagIcon     string
	IsRtl        bool
}

// SiteCode returns code for using in website URLs
func (l Locale) SiteCode() string {
	s := strings.ToLower(l.Code5)
	if s1 := s[:2]; s1 == s[3:] || s1 == "en" || s1 == "fa" || s1 == "ja" || s1 == "zh" {
		return s1
	}
	return s
}

// String represents locale information as string
func (l Locale) String() string {
	return fmt.Sprintf(`Locale{Code5: "%v", IsRtl: %v, NativeTitle: "%v", EnglishTitle: "%v", FlagIcon: "%v"}`, l.Code5, l.IsRtl, l.NativeTitle, l.EnglishTitle, l.FlagIcon)
}

// TitleWithIcon returns name of the language and flag emoji
func (l Locale) TitleWithIcon() string {
	if l.IsRtl {
		return l.NativeTitle + " " + l.FlagIcon
	}
	return l.FlagIcon + " " + l.NativeTitle
}

// TitleWithIconAndNumber returns name, flag emoji and a number // TODO: should bot be here
func (l Locale) TitleWithIconAndNumber(i int) string {
	if l.IsRtl {
		return fmt.Sprintf("%v %v .%d/", l.FlagIcon, l.NativeTitle, i)
	}
	return fmt.Sprintf("/%d. %v %v", i, l.NativeTitle, l.FlagIcon)
}
