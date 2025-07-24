package i18n

import "testing"

func TestLocale_String(t *testing.T) {
	l := Locale{Code5: "c0de5", IsRtl: true, NativeTitle: "Код05", EnglishTitle: "C0de5", FlagIcon: "fl0g"}
	actualLs := l.String()
	expectingLs := `Locale{Code5: "c0de5", IsRtl: true, NativeTitle: "Код05", EnglishTitle: "C0de5", FlagIcon: "fl0g"}`
	if actualLs != expectingLs {
		t.Errorf("Unexpected result of func (Locale) String(). Got: %v. Exepcted: %v", actualLs, expectingLs)
	}
}

func TestLocale_SiteCode(t *testing.T) {
	testCases := []struct {
		name     string
		locale   Locale
		expected string
	}{
		{
			name:     "Same first and last two chars",
			locale:   Locale{Code5: "en-en"},
			expected: "en",
		},
		{
			name:     "English locale",
			locale:   Locale{Code5: "en-US"},
			expected: "en",
		},
		{
			name:     "Farsi locale",
			locale:   Locale{Code5: "fa-IR"},
			expected: "fa",
		},
		{
			name:     "Japanese locale",
			locale:   Locale{Code5: "ja-JP"},
			expected: "ja",
		},
		{
			name:     "Chinese locale",
			locale:   Locale{Code5: "zh-CN"},
			expected: "zh",
		},
		{
			name:     "French locale (first two chars match last two)",
			locale:   Locale{Code5: "fr-FR"},
			expected: "fr",
		},
		{
			name:     "Locale with different first and last two chars",
			locale:   Locale{Code5: "es-MX"},
			expected: "es-mx",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.locale.SiteCode()
			if actual != tc.expected {
				t.Errorf("Expected SiteCode() to return %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestLocale_TitleWithIcon(t *testing.T) {
	testCases := []struct {
		name     string
		locale   Locale
		expected string
	}{
		{
			name:     "LTR locale",
			locale:   Locale{NativeTitle: "English", FlagIcon: "🇺🇸", IsRtl: false},
			expected: "🇺🇸 English",
		},
		{
			name:     "RTL locale",
			locale:   Locale{NativeTitle: "العربية", FlagIcon: "🇸🇦", IsRtl: true},
			expected: "العربية 🇸🇦",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.locale.TitleWithIcon()
			if actual != tc.expected {
				t.Errorf("Expected TitleWithIcon() to return %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestLocale_TitleWithIconAndNumber(t *testing.T) {
	testCases := []struct {
		name     string
		locale   Locale
		number   int
		expected string
	}{
		{
			name:     "LTR locale",
			locale:   Locale{NativeTitle: "English", FlagIcon: "🇺🇸", IsRtl: false},
			number:   1,
			expected: "/1. English 🇺🇸",
		},
		{
			name:     "RTL locale",
			locale:   Locale{NativeTitle: "العربية", FlagIcon: "🇸🇦", IsRtl: true},
			number:   2,
			expected: "🇸🇦 العربية .2/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.locale.TitleWithIconAndNumber(tc.number)
			if actual != tc.expected {
				t.Errorf("Expected TitleWithIconAndNumber(%d) to return %q, got %q", tc.number, tc.expected, actual)
			}
		})
	}
}
