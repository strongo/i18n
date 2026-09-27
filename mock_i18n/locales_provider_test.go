package mock_i18n

import (
	"testing"

	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
)

var _ i18n.LocalesProvider = (*MockLocalesProvider)(nil)

func TestNewMockLocalesProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockLocalesProvider(ctrl)
	if m == nil {
		t.Fatal("localesProvider is nil")
	}

	dummyLocale := i18n.Locale{Code5: "en-US"}
	m.EXPECT().GetLocaleByCode5("en-US").Return(dummyLocale, nil)
	if loc, err := m.GetLocaleByCode5("en-US"); err != nil || loc.Code5 != "en-US" {
		t.Fatalf("unexpected GetLocaleByCode5: %v, %v", loc, err)
	}

	m.EXPECT().SupportedLocales().Return([]i18n.Locale{dummyLocale})
	if locs := m.SupportedLocales(); len(locs) != 1 || locs[0].Code5 != "en-US" {
		t.Fatalf("unexpected SupportedLocales: %v", locs)
	}
}
