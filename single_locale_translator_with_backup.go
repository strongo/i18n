package i18n

import (
	"fmt"
)

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
