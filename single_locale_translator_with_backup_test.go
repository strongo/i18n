package i18n

import (
	"testing"
)

// mockSingleLocaleTranslator is a simple implementation of the SingleLocaleTranslator interface for testing
type mockSingleLocaleTranslator struct {
	locale                 Locale
	translateResult        map[string]string
	translateNoWarnResult  map[string]string
	translateWithMapResult map[string]string
}

func (m mockSingleLocaleTranslator) Locale() Locale {
	return m.locale
}

func (m mockSingleLocaleTranslator) Translate(key string, _ ...any) string {
	if result, ok := m.translateResult[key]; ok {
		return result
	}
	return key // Return key if no translation found
}

func (m mockSingleLocaleTranslator) TranslateNoWarning(key string, _ ...any) string {
	if result, ok := m.translateNoWarnResult[key]; ok {
		return result
	}
	return key // Return key if no translation found
}

func (m mockSingleLocaleTranslator) TranslateWithMap(key string, _ map[string]string) string {
	if result, ok := m.translateWithMapResult[key]; ok {
		return result
	}
	return "" // Return empty string if no translation found
}

func TestNewSingleLocaleTranslatorWithBackup(t *testing.T) {
	// Prepare test data
	primaryLocale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	backupLocale := Locale{Code5: "fr-FR", NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"}

	primary := mockSingleLocaleTranslator{locale: primaryLocale}
	backup := mockSingleLocaleTranslator{locale: backupLocale}

	// Create translator with backup
	translator := NewSingleLocaleTranslatorWithBackup(primary, backup)

	// Verify it works correctly by testing its behavior
	// The Locale method should return the primary locale
	if translator.Locale().Code5 != primaryLocale.Code5 {
		t.Errorf("Expected translator to use primary locale, got %s", translator.Locale().Code5)
	}
}

func TestSingleLocaleTranslatorWithBackup_Locale(t *testing.T) {
	// Prepare test data
	primaryLocale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	backupLocale := Locale{Code5: "fr-FR", NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"}

	primary := mockSingleLocaleTranslator{locale: primaryLocale}
	backup := mockSingleLocaleTranslator{locale: backupLocale}

	// Create translator with backup
	translator := NewSingleLocaleTranslatorWithBackup(primary, backup)

	// Test Locale method - should return primary locale
	locale := translator.Locale()
	if locale.Code5 != primaryLocale.Code5 {
		t.Errorf("Expected Locale() to return primary locale with Code5 %q, got %q",
			primaryLocale.Code5, locale.Code5)
	}
}

func TestSingleLocaleTranslatorWithBackup_Translate(t *testing.T) {
	// Prepare test data
	primaryLocale := Locale{Code5: "en-US"}
	backupLocale := Locale{Code5: "fr-FR"}

	testCases := []struct {
		name           string
		primaryResults map[string]string
		backupResults  map[string]string
		key            string
		args           []any
		expected       string
	}{
		{
			name:           "Primary translation exists",
			primaryResults: map[string]string{"greeting": "Hello"},
			backupResults:  map[string]string{"greeting": "Bonjour"},
			key:            "greeting",
			args:           []any{},
			expected:       "Hello",
		},
		{
			name:           "Primary returns key, backup has translation",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{"greeting": "Bonjour"},
			key:            "greeting",
			args:           []any{},
			expected:       "Bonjour",
		},
		{
			name:           "Primary returns empty, backup has translation",
			primaryResults: map[string]string{"greeting": ""},
			backupResults:  map[string]string{"greeting": "Bonjour"},
			key:            "greeting",
			args:           []any{},
			expected:       "Bonjour",
		},
		{
			name:           "Both return empty string",
			primaryResults: map[string]string{"greeting": ""},
			backupResults:  map[string]string{"greeting": ""},
			key:            "greeting",
			args:           []any{"World"},
			expected:       "greeting(args=[World])",
		},
		{
			name:           "Neither has translation",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{},
			key:            "unknown",
			args:           []any{},
			expected:       "unknown", // Our mock returns the key itself
		},
		{
			name:           "Neither has translation with args",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{},
			key:            "unknown",
			args:           []any{"test"},
			expected:       "unknown", // Our mock returns the key itself
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			primary := mockSingleLocaleTranslator{
				locale:          primaryLocale,
				translateResult: tc.primaryResults,
			}
			backup := mockSingleLocaleTranslator{
				locale:          backupLocale,
				translateResult: tc.backupResults,
			}

			translator := NewSingleLocaleTranslatorWithBackup(primary, backup)

			result := translator.Translate(tc.key, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected Translate(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}

func TestSingleLocaleTranslatorWithBackup_TranslateNoWarning(t *testing.T) {
	// Prepare test data
	primaryLocale := Locale{Code5: "en-US"}
	backupLocale := Locale{Code5: "fr-FR"}

	testCases := []struct {
		name           string
		primaryResults map[string]string
		backupResults  map[string]string
		key            string
		args           []any
		expected       string
	}{
		{
			name:           "Primary translation exists",
			primaryResults: map[string]string{"greeting": "Hello"},
			backupResults:  map[string]string{"greeting": "Bonjour"},
			key:            "greeting",
			args:           []any{},
			expected:       "Hello",
		},
		{
			name:           "Primary returns key, backup has translation",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{"greeting": "Bonjour"},
			key:            "greeting",
			args:           []any{},
			expected:       "Bonjour",
		},
		{
			name:           "Neither has translation",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{},
			key:            "unknown",
			args:           []any{},
			expected:       "unknown", // No formatting with args for TranslateNoWarning
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			primary := mockSingleLocaleTranslator{
				locale:                primaryLocale,
				translateNoWarnResult: tc.primaryResults,
			}
			backup := mockSingleLocaleTranslator{
				locale:                backupLocale,
				translateNoWarnResult: tc.backupResults,
			}

			translator := NewSingleLocaleTranslatorWithBackup(primary, backup)

			result := translator.TranslateNoWarning(tc.key, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected TranslateNoWarning(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}

func TestSingleLocaleTranslatorWithBackup_TranslateWithMap(t *testing.T) {
	// Prepare test data
	primaryLocale := Locale{Code5: "en-US"}
	backupLocale := Locale{Code5: "fr-FR"}

	testCases := []struct {
		name           string
		primaryResults map[string]string
		backupResults  map[string]string
		key            string
		args           map[string]string
		expected       string
	}{
		{
			name:           "Primary translation exists",
			primaryResults: map[string]string{"greeting": "Hello, {name}!"},
			backupResults:  map[string]string{"greeting": "Bonjour, {name}!"},
			key:            "greeting",
			args:           map[string]string{"name": "World"},
			expected:       "Hello, {name}!",
		},
		{
			name:           "Primary returns empty, backup has translation",
			primaryResults: map[string]string{"greeting": ""},
			backupResults:  map[string]string{"greeting": "Bonjour, {name}!"},
			key:            "greeting",
			args:           map[string]string{"name": "World"},
			expected:       "Bonjour, {name}!",
		},
		{
			name:           "Neither has translation",
			primaryResults: map[string]string{},
			backupResults:  map[string]string{},
			key:            "unknown",
			args:           map[string]string{"name": "World"},
			expected:       "", // Empty string when no translation found
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			primary := mockSingleLocaleTranslator{
				locale:                 primaryLocale,
				translateWithMapResult: tc.primaryResults,
			}
			backup := mockSingleLocaleTranslator{
				locale:                 backupLocale,
				translateWithMapResult: tc.backupResults,
			}

			translator := NewSingleLocaleTranslatorWithBackup(primary, backup)

			result := translator.TranslateWithMap(tc.key, tc.args)
			if result != tc.expected {
				t.Errorf("Expected TranslateWithMap(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}
