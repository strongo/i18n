package i18n

import (
	"context"
	"errors"
	"testing"
)

// contextMockTranslator is a simple implementation of the Translator interface for testing
type contextMockTranslator struct {
	locale string
}

func (m contextMockTranslator) Translate(key, locale string, _ ...any) string {
	return key + "_" + locale
}

func (m contextMockTranslator) TranslateWithMap(key, locale string, _ map[string]string) string {
	return key + "_" + locale
}

func (m contextMockTranslator) TranslateNoWarning(key, locale string, _ ...any) string {
	return key + "_" + locale
}

// mockLocalesProvider is a simple implementation of the LocalesProvider interface for testing
type mockLocalesProvider struct {
	locales       []Locale
	shouldFailFor string
}

func (m mockLocalesProvider) SupportedLocales() []Locale {
	return m.locales
}

func (m mockLocalesProvider) GetLocaleByCode5(code5 string) (Locale, error) {
	if code5 == m.shouldFailFor {
		return LocaleUndefined, errors.New("locale not found")
	}

	for _, locale := range m.locales {
		if locale.Code5 == code5 {
			return locale, nil
		}
	}
	return LocaleUndefined, errors.New("locale not found")
}

func TestNewContext(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	locales := []Locale{
		{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"},
		{Code5: "fr-FR", NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"},
	}
	provider := mockLocalesProvider{locales: locales}

	// Create translation context
	translationCtx := NewContext(ctx, provider)

	// Verify it's not nil
	if translationCtx == nil {
		t.Error("Expected NewContext to return a non-nil context")
	}
}

func TestTranslationContext_GetTranslator(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	locales := []Locale{
		{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"},
		{Code5: "fr-FR", NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"},
	}
	provider := mockLocalesProvider{locales: locales}

	// Create translation context
	translationCtx := NewContext(ctx, provider).(*translationContext)

	// Set a mock translator provider
	translationCtx.translatorProvider = func(locale string) Translator {
		return contextMockTranslator{locale: locale}
	}

	// Set a locale
	translationCtx.locale = locales[0]

	// Test GetTranslator
	translator := translationCtx.GetTranslator(ctx)

	// Verify the translator works as expected
	result := translator.Translate("test", "en-US")
	expected := "test_en-US"
	if result != expected {
		t.Errorf("Expected translator to return %q, got %q", expected, result)
	}
}

func TestTranslationContext_SetLocale(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	locales := []Locale{
		{Code5: "en-US", NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"},
		{Code5: "fr-FR", NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"},
	}

	testCases := []struct {
		name          string
		code5         string
		shouldFailFor string
		expectError   bool
	}{
		{
			name:          "Valid locale",
			code5:         "en-US",
			shouldFailFor: "",
			expectError:   false,
		},
		{
			name:          "Invalid locale",
			code5:         "xx-XX",
			shouldFailFor: "",
			expectError:   true,
		},
		{
			name:          "Provider fails",
			code5:         "en-US",
			shouldFailFor: "en-US",
			expectError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			provider := mockLocalesProvider{locales: locales, shouldFailFor: tc.shouldFailFor}

			// Create translation context
			translationCtx := NewContext(ctx, provider).(*translationContext)

			// Test SetLocale
			err := translationCtx.SetLocale(tc.code5)

			if tc.expectError {
				if err == nil {
					t.Error("Expected SetLocale to return an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Verify the locale was set correctly
				if translationCtx.locale.Code5 != tc.code5 {
					t.Errorf("Expected locale to be set to %q, got %q", tc.code5, translationCtx.locale.Code5)
				}
			}
		})
	}
}
