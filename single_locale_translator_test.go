package i18n

import (
	"testing"
)

// mockTranslator is a simple implementation of the Translator interface for testing
type mockTranslator struct {
	translations map[string]map[string]string
}

func (m mockTranslator) Translate(key, locale string, args ...any) string {
	if translations, ok := m.translations[key]; ok {
		if translation, ok := translations[locale]; ok {
			if len(args) > 0 {
				return translation + " " + args[0].(string)
			}
			return translation
		}
	}
	return key
}

func (m mockTranslator) TranslateWithMap(key, locale string, args map[string]string) string {
	if translations, ok := m.translations[key]; ok {
		if translation, ok := translations[locale]; ok {
			result := translation
			for _, v := range args {
				result += " " + v
			}
			return result
		}
	}
	var result = key
	for k := range args {
		result += " $EXTRA(" + k + ")"
	}
	return result
}

func (m mockTranslator) TranslateNoWarning(key, locale string, args ...any) string {
	return m.Translate(key, locale, args...)
}

func TestNewSingleMapTranslator(t *testing.T) {
	// Prepare test data
	locale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	translator := mockTranslator{
		translations: map[string]map[string]string{
			"greeting": {
				"en-US": "Hello",
			},
		},
	}

	// Create single locale translator
	singleTranslator := NewSingleMapTranslator(locale, translator)

	// Verify it's not nil
	if singleTranslator == nil {
		t.Error("Expected NewSingleMapTranslator to return a non-nil translator")
	}
}

func TestSingleLocaleTranslator_Locale(t *testing.T) {
	// Prepare test data
	locale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	translator := mockTranslator{}

	// Create single locale translator
	singleTranslator := NewSingleMapTranslator(locale, translator)

	// Test Locale method
	returnedLocale := singleTranslator.Locale()
	if returnedLocale.Code5 != locale.Code5 {
		t.Errorf("Expected Locale() to return locale with Code5 %q, got %q",
			locale.Code5, returnedLocale.Code5)
	}
	if returnedLocale.NativeTitle != locale.NativeTitle {
		t.Errorf("Expected Locale() to return locale with NativeTitle %q, got %q",
			locale.NativeTitle, returnedLocale.NativeTitle)
	}
}

func TestSingleLocaleTranslator_Translate(t *testing.T) {
	// Prepare test data
	locale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	translator := mockTranslator{
		translations: map[string]map[string]string{
			"greeting": {
				"en-US": "Hello",
			},
			"farewell": {
				"en-US": "Goodbye",
			},
		},
	}

	// Create single locale translator
	singleTranslator := NewSingleMapTranslator(locale, translator)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		args     []any
		expected string
	}{
		{
			name:     "Basic translation",
			key:      "greeting",
			args:     []any{},
			expected: "Hello",
		},
		{
			name:     "Translation with args",
			key:      "greeting",
			args:     []any{"World"},
			expected: "Hello World",
		},
		{
			name:     "Another key",
			key:      "farewell",
			args:     []any{},
			expected: "Goodbye",
		},
		{
			name:     "Key not found",
			key:      "nonexistent",
			args:     []any{},
			expected: "nonexistent", // Should return the key itself
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := singleTranslator.Translate(tc.key, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected Translate(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}

func TestSingleLocaleTranslator_TranslateWithMap(t *testing.T) {
	// Prepare test data
	locale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	translator := mockTranslator{
		translations: map[string]map[string]string{
			"greeting": {
				"en-US": "Hello",
			},
			"farewell": {
				"en-US": "Goodbye",
			},
		},
	}

	// Create single locale translator
	singleTranslator := NewSingleMapTranslator(locale, translator)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		args     map[string]string
		expected string
	}{
		{
			name:     "Basic translation with map",
			key:      "greeting",
			args:     map[string]string{"name": "World"},
			expected: "Hello $EXTRA(name)",
		},
		{
			name:     "Another key with map",
			key:      "farewell",
			args:     map[string]string{"name": "Friend"},
			expected: "Goodbye $EXTRA(name)",
		},
		{
			name:     "Key not found with map",
			key:      "nonexistent",
			args:     map[string]string{"name": "Test"},
			expected: "nonexistent $EXTRA(name)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := singleTranslator.TranslateWithMap(tc.key, tc.args)
			if result != tc.expected {
				t.Errorf("Expected TranslateWithMap(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}

func TestSingleLocaleTranslator_TranslateNoWarning(t *testing.T) {
	// Prepare test data
	locale := Locale{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	translator := mockTranslator{
		translations: map[string]map[string]string{
			"greeting": {
				"en-US": "Hello",
			},
		},
	}

	// Create single locale translator
	singleTranslator := NewSingleMapTranslator(locale, translator)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		args     []any
		expected string
	}{
		{
			name:     "Existing translation",
			key:      "greeting",
			args:     []any{"World"},
			expected: "Hello World",
		},
		{
			name:     "Non-existent key",
			key:      "nonexistent",
			args:     []any{},
			expected: "nonexistent", // Should return the key itself
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// This test mainly verifies that the method doesn't panic
			// We can't easily test that no warning is logged
			result := singleTranslator.TranslateNoWarning(tc.key, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected TranslateNoWarning(%q, %v) to return %q, got %q",
					tc.key, tc.args, tc.expected, result)
			}
		})
	}
}
