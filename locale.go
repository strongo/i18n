package i18n

import (
	"fmt"
	"strings"
)

// Locale describes language
type Locale struct {
	Code5        string
	NativeTitle  string
	EnglishTitle string
	FlagIcon     string
	IsRtl        bool
}

// SiteCode returns code for using in website URLs
func (l Locale) SiteCode() string {
	s := strings.ToLower(l.Code5)
	if s1 := s[:2]; s1 == s[3:] || s1 == "en" || s1 == "fa" || s1 == "ja" || s1 == "zh" {
		return s1
	}
	return s
}

// String represents locale information as string
func (l Locale) String() string {
	return fmt.Sprintf(`Locale{Code5: "%v", IsRtl: %v, NativeTitle: "%v", EnglishTitle: "%v", FlagIcon: "%v"}`, l.Code5, l.IsRtl, l.NativeTitle, l.EnglishTitle, l.FlagIcon)
}

// TitleWithIcon returns name of the language and flag emoji
func (l Locale) TitleWithIcon() string {
	if l.IsRtl {
		return l.NativeTitle + " " + l.FlagIcon
	}
	return l.FlagIcon + " " + l.NativeTitle
}

// TitleWithIconAndNumber returns name, flag emoji and a number // TODO: should bot be here
func (l Locale) TitleWithIconAndNumber(i int) string {
	if l.IsRtl {
		return fmt.Sprintf("%v %v .%d/", l.FlagIcon, l.NativeTitle, i)
	}
	return fmt.Sprintf("/%d. %v %v", i, l.NativeTitle, l.FlagIcon)
}
