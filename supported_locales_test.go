package i18n

import (
	"testing"
)

func TestSupported_GetLocaleByCode5(t *testing.T) {
	// Create a supported locales provider with a few locales
	locales := []Locale{
		LocaleEnUS,
		LocaleRuRU,
		LocaleFrFR,
	}
	provider := supported{locales: locales}

	testCases := []struct {
		name          string
		code5         string
		expectedCode5 string
		expectError   bool
	}{
		{
			name:          "Full code - en-US",
			code5:         "en-US",
			expectedCode5: "en-US",
			expectError:   false,
		},
		{
			name:          "Short code - en",
			code5:         "en",
			expectedCode5: "en-US",
			expectError:   false,
		},
		{
			name:          "Full code - ru-RU",
			code5:         "ru-RU",
			expectedCode5: "ru-RU",
			expectError:   false,
		},
		{
			name:          "Short code - ru",
			code5:         "ru",
			expectedCode5: "ru-RU",
			expectError:   false,
		},
		{
			name:        "Unknown locale",
			code5:       "xx-XX",
			expectError: true,
		},
		{
			name:        "Empty code",
			code5:       "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			locale, err := provider.GetLocaleByCode5(tc.code5)

			if tc.expectError {
				if err == nil {
					t.Error("Expected an error, but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if locale.Code5 != tc.expectedCode5 {
				t.Errorf("Expected locale code to be %q, got %q", tc.expectedCode5, locale.Code5)
			}
		})
	}
}

func TestSupported_SupportedLocales(t *testing.T) {
	testCases := []struct {
		name          string
		locales       []Locale
		expectedCount int
	}{
		{
			name:          "Empty list",
			locales:       []Locale{},
			expectedCount: 0,
		},
		{
			name:          "Single locale",
			locales:       []Locale{LocaleEnUS},
			expectedCount: 1,
		},
		{
			name:          "Multiple locales",
			locales:       []Locale{LocaleEnUS, LocaleRuRU, LocaleFrFR},
			expectedCount: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the original locales for comparison
			originalLocales := make([]Locale, len(tc.locales))
			copy(originalLocales, tc.locales)

			provider := supported{locales: tc.locales}
			result := provider.SupportedLocales()

			// Check count
			if len(result) != tc.expectedCount {
				t.Errorf("Expected %d locales, got %d", tc.expectedCount, len(result))
			}

			// Check that all locales are included with correct values
			for i, locale := range originalLocales {
				if i < len(result) {
					if result[i].Code5 != locale.Code5 {
						t.Errorf("Expected locale at index %d to have code %q, got %q",
							i, locale.Code5, result[i].Code5)
					}
				}
			}

			// Test that the returned slice is a copy by modifying it
			if len(result) > 0 {
				// Save the original value
				originalCode5 := result[0].Code5

				// Modify the returned slice
				result[0].Code5 = "modified"

				// Check that the original slice is not affected
				if len(tc.locales) > 0 && tc.locales[0].Code5 != originalCode5 {
					t.Errorf("Expected original locale code to remain %q, got %q. The returned slice should be a copy.",
						originalCode5, tc.locales[0].Code5)
				}
			}
		})
	}
}
