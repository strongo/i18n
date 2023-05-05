package i18n

import "context"

// TranslationContext is an i18n context (internationalization)
type TranslationContext interface {
	GetTranslator(c context.Context) SingleLocaleTranslator
	SupportedLocales() LocalesProvider
	SetLocale(code5 string) error
}

type translationContext struct {
	ctx                context.Context
	supportedLocales   LocalesProvider
	locale             Locale
	translatorProvider TranslatorProvider
}

func (l10n translationContext) GetTranslator(c context.Context) SingleLocaleTranslator {
	return l10n.translatorProvider(l10n.locale.Code5)
}

func (l10n translationContext) SupportedLocales() LocalesProvider {
	return l10n.supportedLocales
}

func NewContext(c context.Context, supportedLocales LocalesProvider) TranslationContext {
	return translationContext{ctx: c, supportedLocales: supportedLocales}
}

func (l10n translationContext) SetLocale(code5 string) error {
	locale, err := l10n.supportedLocales.GetLocaleByCode5(code5)
	if err != nil {
		errorf(l10n.ctx, "*WebhookContextBase.SetLocate(%v) - %v", code5, err)
		return err
	}
	l10n.locale = locale
	debugf(l10n.ctx, "*WebhookContextBase.SetLocale(%v) => Done", code5)
	return nil
}
