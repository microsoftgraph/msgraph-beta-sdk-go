package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RegionalAndLanguageSettingsable 
type RegionalAndLanguageSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthoringLanguages()([]LocaleInfoable)
    GetDefaultDisplayLanguage()(LocaleInfoable)
    GetDefaultRegionalFormat()(LocaleInfoable)
    GetDefaultSpeechInputLanguage()(LocaleInfoable)
    GetDefaultTranslationLanguage()(LocaleInfoable)
    GetRegionalFormatOverrides()(RegionalFormatOverridesable)
    GetTranslationPreferences()(TranslationPreferencesable)
    SetAuthoringLanguages(value []LocaleInfoable)()
    SetDefaultDisplayLanguage(value LocaleInfoable)()
    SetDefaultRegionalFormat(value LocaleInfoable)()
    SetDefaultSpeechInputLanguage(value LocaleInfoable)()
    SetDefaultTranslationLanguage(value LocaleInfoable)()
    SetRegionalFormatOverrides(value RegionalFormatOverridesable)()
    SetTranslationPreferences(value TranslationPreferencesable)()
}
