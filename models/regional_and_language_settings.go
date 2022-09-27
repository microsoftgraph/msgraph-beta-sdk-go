package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegionalAndLanguageSettings 
type RegionalAndLanguageSettings struct {
    Entity
    // Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
    authoringLanguages []LocaleInfoable
    // The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
    defaultDisplayLanguage LocaleInfoable
    // The locale that drives the default date, time, and calendar formatting.Returned by default.
    defaultRegionalFormat LocaleInfoable
    // The language a user expected to use as input for text to speech scenarios.Returned by default.
    defaultSpeechInputLanguage LocaleInfoable
    // The language a user expects to have documents, emails, and messages translated into.Returned by default.
    defaultTranslationLanguage LocaleInfoable
    // Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
    regionalFormatOverrides RegionalFormatOverridesable
    // The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
    translationPreferences TranslationPreferencesable
}
// NewRegionalAndLanguageSettings instantiates a new regionalAndLanguageSettings and sets the default values.
func NewRegionalAndLanguageSettings()(*RegionalAndLanguageSettings) {
    m := &RegionalAndLanguageSettings{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.regionalAndLanguageSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRegionalAndLanguageSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegionalAndLanguageSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegionalAndLanguageSettings(), nil
}
// GetAuthoringLanguages gets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetAuthoringLanguages()([]LocaleInfoable) {
    return m.authoringLanguages
}
// GetDefaultDisplayLanguage gets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetDefaultDisplayLanguage()(LocaleInfoable) {
    return m.defaultDisplayLanguage
}
// GetDefaultRegionalFormat gets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultRegionalFormat()(LocaleInfoable) {
    return m.defaultRegionalFormat
}
// GetDefaultSpeechInputLanguage gets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultSpeechInputLanguage()(LocaleInfoable) {
    return m.defaultSpeechInputLanguage
}
// GetDefaultTranslationLanguage gets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultTranslationLanguage()(LocaleInfoable) {
    return m.defaultTranslationLanguage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegionalAndLanguageSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authoringLanguages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLocaleInfoFromDiscriminatorValue , m.SetAuthoringLanguages)
    res["defaultDisplayLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocaleInfoFromDiscriminatorValue , m.SetDefaultDisplayLanguage)
    res["defaultRegionalFormat"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocaleInfoFromDiscriminatorValue , m.SetDefaultRegionalFormat)
    res["defaultSpeechInputLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocaleInfoFromDiscriminatorValue , m.SetDefaultSpeechInputLanguage)
    res["defaultTranslationLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocaleInfoFromDiscriminatorValue , m.SetDefaultTranslationLanguage)
    res["regionalFormatOverrides"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRegionalFormatOverridesFromDiscriminatorValue , m.SetRegionalFormatOverrides)
    res["translationPreferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTranslationPreferencesFromDiscriminatorValue , m.SetTranslationPreferences)
    return res
}
// GetRegionalFormatOverrides gets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) GetRegionalFormatOverrides()(RegionalFormatOverridesable) {
    return m.regionalFormatOverrides
}
// GetTranslationPreferences gets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetTranslationPreferences()(TranslationPreferencesable) {
    return m.translationPreferences
}
// Serialize serializes information the current object
func (m *RegionalAndLanguageSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthoringLanguages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthoringLanguages())
        err = writer.WriteCollectionOfObjectValues("authoringLanguages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultDisplayLanguage", m.GetDefaultDisplayLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultRegionalFormat", m.GetDefaultRegionalFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultSpeechInputLanguage", m.GetDefaultSpeechInputLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultTranslationLanguage", m.GetDefaultTranslationLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("regionalFormatOverrides", m.GetRegionalFormatOverrides())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("translationPreferences", m.GetTranslationPreferences())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthoringLanguages sets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetAuthoringLanguages(value []LocaleInfoable)() {
    m.authoringLanguages = value
}
// SetDefaultDisplayLanguage sets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetDefaultDisplayLanguage(value LocaleInfoable)() {
    m.defaultDisplayLanguage = value
}
// SetDefaultRegionalFormat sets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultRegionalFormat(value LocaleInfoable)() {
    m.defaultRegionalFormat = value
}
// SetDefaultSpeechInputLanguage sets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultSpeechInputLanguage(value LocaleInfoable)() {
    m.defaultSpeechInputLanguage = value
}
// SetDefaultTranslationLanguage sets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultTranslationLanguage(value LocaleInfoable)() {
    m.defaultTranslationLanguage = value
}
// SetRegionalFormatOverrides sets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) SetRegionalFormatOverrides(value RegionalFormatOverridesable)() {
    m.regionalFormatOverrides = value
}
// SetTranslationPreferences sets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetTranslationPreferences(value TranslationPreferencesable)() {
    m.translationPreferences = value
}
