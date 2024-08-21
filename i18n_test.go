package i18n

import (
	"context"
	"testing"
)

func TestLocale_String(t *testing.T) {
	l := Locale{Code5: "c0de5", IsRtl: true, NativeTitle: "Код05", EnglishTitle: "C0de5", FlagIcon: "fl0g"}
	actualLs := l.String()
	expectingLs := `Locale{Code5: "c0de5", IsRtl: true, NativeTitle: "Код05", EnglishTitle: "C0de5", FlagIcon: "fl0g"}`
	if actualLs != expectingLs {
		t.Errorf("Unexpected result of func (Locale) String(). Got: %v. Exepcted: %v", actualLs, expectingLs)
	}
}

func TestNewSingleLocaleTranslatorWithBackup(t *testing.T) {
	type args struct {
		primary SingleLocaleTranslator
		backup  SingleLocaleTranslator
	}
	ctx := context.Background()
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "has_both_primary_and_backup",
			args: args{
				primary: NewSingleMapTranslator(LocaleEnUS, NewMapTranslator(ctx, map[string]map[string]string{"s1": {"en-US": "United States"}})),
				backup:  NewSingleMapTranslator(LocaleEnUK, NewMapTranslator(ctx, map[string]map[string]string{"s1": {"en-UK": "United Kingdom"}})),
			},
			want: "United States",
		},
		{
			name: "has_only_primary",
			args: args{
				primary: NewSingleMapTranslator(LocaleEnUS, NewMapTranslator(ctx, map[string]map[string]string{"s1": {"en-US": "United States"}})),
				backup:  NewSingleMapTranslator(LocaleEnUK, NewMapTranslator(ctx, map[string]map[string]string{"s2": {"en-UK": "United Kingdom"}})),
			},
			want: "United States",
		},
		{
			name: "has_only_backup",
			args: args{
				primary: NewSingleMapTranslator(LocaleEnUS, NewMapTranslator(ctx, map[string]map[string]string{"s2": {"en-US": "United States"}})),
				backup:  NewSingleMapTranslator(LocaleEnUK, NewMapTranslator(ctx, map[string]map[string]string{"s1": {"en-UK": "United Kingdom"}})),
			},
			want: "United Kingdom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := NewSingleLocaleTranslatorWithBackup(tt.args.primary, tt.args.backup)

			if got := translator.Translate("s1"); got != tt.want {
				t.Errorf("NewSingleLocaleTranslatorWithBackup().Transalge() = %s, want %s", got, tt.want)
			}
		})
	}
}
