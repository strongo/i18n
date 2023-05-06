package i18n

import (
	"context"
	"testing"
)

func TestNewContext(t *testing.T) {
	NewContext(context.Background(), NewSupportedLocales([]string{"en-US", "ua-UA"}))
}
