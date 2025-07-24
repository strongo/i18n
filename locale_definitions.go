package i18n

import (
	"fmt"
)

// TODO: This module should be in a dedicate package?

//"4. French ",
//"5. Spanish ",
//"6. Italian \xF0\x9F\x87\xAE\xF0\x9F\x87\xB9",

var LocaleUndefined = Locale{Code5: LocaleCodeUndefined, NativeTitle: "Undefined", EnglishTitle: "Undefined"}

var (
	LocaleArEG = Locale{Code5: LocaleCodeArEG, NativeTitle: "العربية المصرية", EnglishTitle: "Arabic Egypt", FlagIcon: "🇪🇬"}
	LocaleDeDE = Locale{Code5: LocaleCodeDeDE, NativeTitle: "Deutsch", EnglishTitle: "German", FlagIcon: "🇩🇪"}
	LocaleEnUK = Locale{Code5: LocaleCodeEnUK, NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇬🇧"}
	LocaleEnUS = Locale{Code5: LocaleCodeEnUS, NativeTitle: "English", EnglishTitle: "English", FlagIcon: "🇺🇸"}
	LocaleEsES = Locale{Code5: LocaleCodeEsES, NativeTitle: "Español", EnglishTitle: "Spanish", FlagIcon: "🇪🇸"}
	LocaleFaIR = Locale{Code5: LocaleCodeFaIR, IsRtl: true, NativeTitle: "فارسی", EnglishTitle: "Farsi", FlagIcon: "🇮🇷"}
	LocaleFrFR = Locale{Code5: LocaleCodeFrFR, NativeTitle: "Français", EnglishTitle: "French", FlagIcon: "🇫🇷"}
	LocaleIdID = Locale{Code5: LocaleCodeIdID, NativeTitle: "Bahasa Indonesia", EnglishTitle: "Indonesian", FlagIcon: "🇮🇩"}
	LocaleItIT = Locale{Code5: LocaleCodeItIT, NativeTitle: "Italiano", EnglishTitle: "Italian", FlagIcon: "🇮🇹"}
	LocaleJaJP = Locale{Code5: LocaleCodeJaJP, NativeTitle: "日本語", EnglishTitle: "Japanese", FlagIcon: "🇯🇵"}
	LocaleKoKR = Locale{Code5: LocaleCodeKoKR, NativeTitle: "한국어", EnglishTitle: "Korean", FlagIcon: "🇰🇷"}
	LocalePlPl = Locale{Code5: LocaleCodePlPL, NativeTitle: "Polszczyzna", EnglishTitle: "Polish", FlagIcon: "🇵🇱"}
	LocalePtBR = Locale{Code5: LocaleCodePtBR, NativeTitle: "Português (BR)", EnglishTitle: "Portuguese (BR)", FlagIcon: "🇧🇷"}
	LocalePtPT = Locale{Code5: LocaleCodePtPT, NativeTitle: "Português (PT)", EnglishTitle: "Portuguese (PT)", FlagIcon: "🇵🇹"}
	LocaleRuRU = Locale{Code5: LocaleCodeRuRU, NativeTitle: "Русский", EnglishTitle: "Russian", FlagIcon: "🇷🇺"}
	LocaleTrTR = Locale{Code5: LocaleCodeTrTR, NativeTitle: "Türkçe", EnglishTitle: "Turkish", FlagIcon: "🇹🇷"}
	LocaleUkUA = Locale{Code5: LocaleCodeUkUA, NativeTitle: "Українська", EnglishTitle: "Ukrainian", FlagIcon: "🇺🇦"}
	LocaleUzUZ = Locale{Code5: LocaleCodeUzUZ, NativeTitle: "Oʻzbek tili", EnglishTitle: "Uzbek", FlagIcon: "🇨🇳"}
	LocaleZhCN = Locale{Code5: LocaleCodeZhCN, NativeTitle: "中文", EnglishTitle: "Chinese", FlagIcon: "🇨🇳"}
)

// LocalesByCode5 map of locales by 5-character code
var LocalesByCode5 = map[string]Locale{
	LocaleCodeArEG: LocaleArEG,
	LocaleCodeDeDE: LocaleDeDE,
	LocaleCodeEnUK: LocaleEnUK,
	LocaleCodeEnUS: LocaleEnUS,
	LocaleCodeEsES: LocaleEsES,
	LocaleCodeFaIR: LocaleFaIR,
	LocaleCodeFrFR: LocaleFrFR,
	LocaleCodeIdID: LocaleIdID,
	LocaleCodeItIT: LocaleItIT,
	LocaleCodeJaJP: LocaleJaJP,
	LocaleCodeKoKR: LocaleKoKR,
	LocaleCodePlPL: LocalePlPl,
	LocaleCodePtBR: LocalePtBR,
	LocaleCodePtPT: LocalePtPT,
	LocaleCodeRuRU: LocaleRuRU,
	LocaleCodeTrTR: LocaleTrTR,
	LocaleCodeUkUA: LocaleUkUA,
	LocaleCodeUzUZ: LocaleUzUZ,
	LocaleCodeZhCN: LocaleZhCN,
}

// GetLocaleByCode5 returns locale by 5-character code
func GetLocaleByCode5(code5 string) Locale {
	if locale, ok := LocalesByCode5[code5]; ok {
		return locale
	}
	panic(fmt.Sprintf("Unknown locale: [%v]", code5))
}

func NewSupportedLocales(code5s []string) LocalesProvider {
	s := supported{
		locales: make([]Locale, len(code5s)),
	}
	for i, code5 := range code5s {
		s.locales[i] = GetLocaleByCode5(code5)
	}
	return s
}
