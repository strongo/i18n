package i18n

import "context"

// TranslationContext is an i18n context (internationalization)
type TranslationContext interface {
	LocalesProvider
	GetTranslator(c context.Context) Translator
	SetLocale(code5 string) error
}

func NewContext(c context.Context, supportedLocales LocalesProvider) TranslationContext {
	return &translationContext{ctx: c, LocalesProvider: supportedLocales}
}

type translationContext struct {
	ctx                context.Context
	locale             Locale
	translatorProvider TranslatorProvider
	LocalesProvider
}

func (l10n *translationContext) GetTranslator(_ context.Context) Translator {
	return l10n.translatorProvider(l10n.locale.Code5)
}

func (l10n *translationContext) SetLocale(code5 string) error {
	locale, err := l10n.GetLocaleByCode5(code5)
	if err != nil {
		errorf(l10n.ctx, "*WebhookContextBase.SetLocate(%v) - %v", code5, err)
		return err
	}
	l10n.locale = locale
	debugf(l10n.ctx, "*WebhookContextBase.SetLocale(%v) => Done", code5)
	return nil
}
