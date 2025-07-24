package i18n

import (
	"testing"
)

func TestGetLocaleByCode5(t *testing.T) {
	testCases := []struct {
		name          string
		code5         string
		expectedCode5 string
		shouldPanic   bool
	}{
		{
			name:          "English US",
			code5:         LocaleCodeEnUS,
			expectedCode5: LocaleCodeEnUS,
			shouldPanic:   false,
		},
		{
			name:          "Russian",
			code5:         LocaleCodeRuRU,
			expectedCode5: LocaleCodeRuRU,
			shouldPanic:   false,
		},
		{
			name:        "Unknown locale",
			code5:       "xx-XX",
			shouldPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Expected GetLocaleByCode5 to panic for unknown locale, but it didn't")
					}
				}()
			}

			locale := GetLocaleByCode5(tc.code5)
			if !tc.shouldPanic && locale.Code5 != tc.expectedCode5 {
				t.Errorf("Expected locale code to be %q, got %q", tc.expectedCode5, locale.Code5)
			}
		})
	}
}

func TestNewSupportedLocales(t *testing.T) {
	testCases := []struct {
		name          string
		code5s        []string
		expectedCount int
		expectedCodes []string
		shouldPanic   bool
	}{
		{
			name:          "Empty list",
			code5s:        []string{},
			expectedCount: 0,
			expectedCodes: []string{},
			shouldPanic:   false,
		},
		{
			name:          "Single locale",
			code5s:        []string{LocaleCodeEnUS},
			expectedCount: 1,
			expectedCodes: []string{LocaleCodeEnUS},
			shouldPanic:   false,
		},
		{
			name:          "Multiple locales",
			code5s:        []string{LocaleCodeEnUS, LocaleCodeRuRU, LocaleCodeFrFR},
			expectedCount: 3,
			expectedCodes: []string{LocaleCodeEnUS, LocaleCodeRuRU, LocaleCodeFrFR},
			shouldPanic:   false,
		},
		{
			name:        "Unknown locale",
			code5s:      []string{LocaleCodeEnUS, "xx-XX"},
			shouldPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Expected NewSupportedLocales to panic for unknown locale, but it didn't")
					}
				}()
			}

			provider := NewSupportedLocales(tc.code5s)

			if !tc.shouldPanic {
				locales := provider.SupportedLocales()
				if len(locales) != tc.expectedCount {
					t.Errorf("Expected %d locales, got %d", tc.expectedCount, len(locales))
				}

				for i, expectedCode := range tc.expectedCodes {
					if i < len(locales) && locales[i].Code5 != expectedCode {
						t.Errorf("Expected locale at index %d to have code %q, got %q",
							i, expectedCode, locales[i].Code5)
					}
				}
			}
		})
	}
}

func TestLocalesByCode5(t *testing.T) {
	// Test that all locale codes have corresponding entries in the map
	for code := range LocalesByCode5 {
		if _, ok := LocalesByCode5[code]; !ok {
			t.Errorf("Locale code %q is not in LocalesByCode5 map", code)
		}
	}

	// Test a few specific locales
	testCases := []struct {
		code5         string
		expectedTitle string
	}{
		{LocaleCodeEnUS, "English"},
		{LocaleCodeRuRU, "Русский"},
		{LocaleCodeFrFR, "Français"},
	}

	for _, tc := range testCases {
		locale, ok := LocalesByCode5[tc.code5]
		if !ok {
			t.Errorf("Locale code %q not found in LocalesByCode5", tc.code5)
			continue
		}
		if locale.NativeTitle != tc.expectedTitle {
			t.Errorf("Expected locale %q to have title %q, got %q",
				tc.code5, tc.expectedTitle, locale.NativeTitle)
		}
	}
}
