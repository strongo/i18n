package i18n

import (
	"context"
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx := context.Background()
	supportedLocales := NewSupportedLocales([]string{
		"ar-EG",
		"en-US",
		"ja-JP",
		"ko-KR",
		"ru-RU",
		"uk-UA",
		"zh-CN",
	})
	tc := NewContext(ctx, supportedLocales)
	if tc == nil {
		t.Errorf("func NewContext() returned nil")
	}
}
