package i18n

import (
	"context"
	"strings"
	"testing"
)

func TestNewMapTranslator(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	defaultLocale := "en-US"
	translations := map[string]map[string]string{
		"greeting": {
			"en-US": "Hello",
			"es-ES": "Hola",
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Verify it's not nil
	if translator == nil {
		t.Error("Expected NewMapTranslator to return a non-nil translator")
	}
}

func TestMapTranslator_Translate(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	defaultLocale := "en-US"
	translations := map[string]map[string]string{
		"greeting": {
			"en-US": "Hello, %s!",
			"es-ES": "¡Hola, %s!",
		},
		"welcome": {
			"en-US": "Welcome to our website",
			// No Spanish translation for this key
		},
		"template": {
			"en-US": "Hello, {{.Name}}!",
		},
		"template_with_error": {
			"en-US": "Hello, {{if .NonExistentMethod}}{{.Name}}{{end}}!",
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		locale   string
		args     []any
		expected string
	}{
		{
			name:     "Basic translation",
			key:      "greeting",
			locale:   "en-US",
			args:     []any{"World"},
			expected: "Hello, World!",
		},
		{
			name:     "Translation in another locale",
			key:      "greeting",
			locale:   "es-ES",
			args:     []any{"Mundo"},
			expected: "¡Hola, Mundo!",
		},
		{
			name:     "Fallback to default locale",
			key:      "welcome",
			locale:   "es-ES", // No Spanish translation
			args:     []any{},
			expected: "Welcome to our website",
		},
		{
			name:     "Key not found",
			key:      "nonexistent",
			locale:   "en-US",
			args:     []any{},
			expected: "", // The actual implementation returns an empty string for non-existent keys
		},
		{
			name:     "Template with struct",
			key:      "template",
			locale:   "en-US",
			args:     []any{struct{ Name string }{"World"}},
			expected: "Hello, World!",
		},
		{
			name:     "Template used multiple times",
			key:      "template",
			locale:   "en-US",
			args:     []any{struct{ Name string }{"Universe"}},
			expected: "Hello, Universe!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := translator.Translate(tc.key, tc.locale, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected Translate(%q, %q, %v) to return %q, got %q",
					tc.key, tc.locale, tc.args, tc.expected, result)
			}
		})
	}
}

func TestMapTranslator_TranslateWithMap(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	defaultLocale := "en-US"
	translations := map[string]map[string]string{
		"greeting": {
			"en-US": "Hello, {name}!",
			"es-ES": "¡Hola, {name}!",
		},
		"profile": {
			"en-US": "Name: {name}, Age: {age}",
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		locale   string
		args     map[string]string
		expected string
	}{
		{
			name:     "Basic translation with map",
			key:      "greeting",
			locale:   "en-US",
			args:     map[string]string{"name": "World"},
			expected: "Hello, World!",
		},
		{
			name:     "Translation in another locale with map",
			key:      "greeting",
			locale:   "es-ES",
			args:     map[string]string{"name": "Mundo"},
			expected: "¡Hola, Mundo!",
		},
		{
			name:     "Multiple placeholders",
			key:      "profile",
			locale:   "en-US",
			args:     map[string]string{"name": "John", "age": "30"},
			expected: "Name: John, Age: 30",
		},
		{
			name:     "Extra arguments",
			key:      "greeting",
			locale:   "en-US",
			args:     map[string]string{"name": "World", "extra": "value"},
			expected: "Hello, World! $EXTRA(extra)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := translator.TranslateWithMap(tc.key, tc.locale, tc.args)
			if result != tc.expected {
				t.Errorf("Expected TranslateWithMap(%q, %q, %v) to return %q, got %q",
					tc.key, tc.locale, tc.args, tc.expected, result)
			}
		})
	}
}

func TestMapTranslator_TranslateNoWarning(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	defaultLocale := "en-US"
	translations := map[string]map[string]string{
		"greeting": {
			"en-US": "Hello, %s!",
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Test cases
	testCases := []struct {
		name     string
		key      string
		locale   string
		args     []any
		expected string
	}{
		{
			name:     "Existing translation",
			key:      "greeting",
			locale:   "en-US",
			args:     []any{"World"},
			expected: "Hello, World!",
		},
		{
			name:     "Non-existent key",
			key:      "nonexistent",
			locale:   "en-US",
			args:     []any{},
			expected: "", // The actual implementation returns an empty string for non-existent keys
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// This test mainly verifies that the method doesn't panic
			// We can't easily test that no warning is logged
			result := translator.TranslateNoWarning(tc.key, tc.locale, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected TranslateNoWarning(%q, %q, %v) to return %q, got %q",
					tc.key, tc.locale, tc.args, tc.expected, result)
			}
		})
	}
}

func TestPlaceMapValues(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		args     map[string]string
		expected string
	}{
		{
			name:     "Single placeholder",
			s:        "Hello, {name}!",
			args:     map[string]string{"name": "World"},
			expected: "Hello, World!",
		},
		{
			name:     "Multiple placeholders",
			s:        "Name: {name}, Age: {age}",
			args:     map[string]string{"name": "John", "age": "30"},
			expected: "Name: John, Age: 30",
		},
		{
			name:     "Repeated placeholders",
			s:        "{name} is {name}",
			args:     map[string]string{"name": "John"},
			expected: "John is John",
		},
		{
			name:     "Extra arguments",
			s:        "Hello, {name}!",
			args:     map[string]string{"name": "World", "extra": "value"},
			expected: "Hello, World! $EXTRA(extra)",
		},
		{
			name:     "No placeholders",
			s:        "Hello, World!",
			args:     map[string]string{"name": "John"},
			expected: "Hello, World! $EXTRA(name)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := placeMapValues(tc.s, tc.args)
			if result != tc.expected {
				t.Errorf("Expected placeMapValues(%q, %v) to return %q, got %q",
					tc.s, tc.args, tc.expected, result)
			}
		})
	}
}

func TestMapTranslator_EmptyDefaultLocale(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	// Empty default locale should fall back to "en-US"
	defaultLocale := ""
	translations := map[string]map[string]string{
		"greeting": {
			"en-US": "Hello!",
			"es-ES": "¡Hola!",
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Test fallback to en-US when default locale is empty
	result := translator.Translate("greeting", "fr-FR") // Not in translations
	expected := "Hello!"                                // Should fall back to en-US
	if result != expected {
		t.Errorf("Expected Translate(%q, %q) to return %q (fallback to en-US), got %q",
			"greeting", "fr-FR", expected, result)
	}
}

func TestMapTranslator_DefaultLocaleNotFound(t *testing.T) {
	// Prepare test data
	ctx := context.Background()
	defaultLocale := "en-US"
	translations := map[string]map[string]string{
		"greeting": {
			"es-ES": "¡Hola!",
			// No translation for en-US
		},
	}

	// Create translator
	translator := NewMapTranslator(ctx, defaultLocale, translations)

	// Test key is returned when translation not found for default locale
	result := translator.Translate("greeting", "fr-FR") // Not in translations
	expected := "greeting"                              // Should return the key itself
	if result != expected {
		t.Errorf("Expected Translate(%q, %q) to return %q (key itself), got %q",
			"greeting", "fr-FR", expected, result)
	}

	// Test with TranslateNoWarning
	result = translator.TranslateNoWarning("greeting", "fr-FR") // Not in translations
	if result != expected {
		t.Errorf("Expected TranslateNoWarning(%q, %q) to return %q (key itself), got %q",
			"greeting", "fr-FR", expected, result)
	}
}

func TestMapTranslator_TemplatePanics(t *testing.T) {
	// Test case 1: Template parsing error
	t.Run("Template parsing error", func(t *testing.T) {
		// Prepare test data with invalid template syntax
		ctx := context.Background()
		defaultLocale := "en-US"
		translations := map[string]map[string]string{
			"invalid_template": {
				// This template has invalid syntax but includes the markers that trigger template parsing:
				// 1. Contains "}}" - closing template marker
				// 2. Contains "{{." - opening template marker with dot
				"en-US": "Hello, {{.Name}} and {{.Invalid}",
			},
		}

		// Create translator
		translator := NewMapTranslator(ctx, defaultLocale, translations).(mapTranslator)

		defer func() {
			r := recover()
			if r == nil {
				t.Error("Expected a panic when parsing an invalid template, but got none")
				return
			}

			// Verify the panic message contains expected text
			panicMsg, ok := r.(string)
			if !ok {
				t.Errorf("Expected panic value to be a string, got %T", r)
				return
			}
			if !strings.Contains(panicMsg, "Failed to parse template") {
				t.Errorf("Expected panic message to contain 'Failed to parse template', got: %s", panicMsg)
			}
		}()

		// This should trigger a panic in template parsing
		// We need to:
		// 1. Provide exactly one argument (to satisfy len(args) == 1)
		// 2. Use a template with both "{{." and "}}" to trigger template parsing
		translator._translate(true, "invalid_template", "en-US", struct{ Name string }{"World"})
	})

	// Test case 2: Template execution error
	t.Run("Template execution error", func(t *testing.T) {
		// Prepare test data with template that will cause execution error
		ctx := context.Background()
		defaultLocale := "en-US"
		translations := map[string]map[string]string{
			"template_with_error": {
				// This template has valid syntax but will fail during execution
				// because it tries to call a method that doesn't exist
				"en-US": "Hello, {{.Name}} and {{.NonExistentMethod}}!",
			},
		}

		// Create translator
		translator := NewMapTranslator(ctx, defaultLocale, translations).(mapTranslator)

		defer func() {
			r := recover()
			if r == nil {
				t.Error("Expected a panic when executing a template with errors, but got none")
				return
			}

			// Verify the panic message contains expected text
			panicMsg, ok := r.(string)
			if !ok {
				t.Errorf("Expected panic value to be a string, got %T", r)
				return
			}
			if !strings.Contains(panicMsg, "Failed to render template") {
				t.Errorf("Expected panic message to contain 'Failed to render template', got: %s", panicMsg)
			}
		}()

		// This should trigger a panic in template execution
		// We need to:
		// 1. Provide exactly one argument (to satisfy len(args) == 1)
		// 2. Use a template with both "{{." and "}}" to trigger template execution
		// 3. The template must be valid syntax but cause an execution error
		translator._translate(true, "template_with_error", "en-US", struct{ Name string }{"World"})
	})
}
