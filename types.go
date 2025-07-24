package i18n

type TranslatorProvider = func(locale string) Translator
