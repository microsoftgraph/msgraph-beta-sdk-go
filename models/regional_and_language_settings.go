package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegionalAndLanguageSettings 
type RegionalAndLanguageSettings struct {
    Entity
}
// NewRegionalAndLanguageSettings instantiates a new regionalAndLanguageSettings and sets the default values.
func NewRegionalAndLanguageSettings()(*RegionalAndLanguageSettings) {
    m := &RegionalAndLanguageSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRegionalAndLanguageSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegionalAndLanguageSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegionalAndLanguageSettings(), nil
}
// GetAuthoringLanguages gets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetAuthoringLanguages()([]LocaleInfoable) {
    val, err := m.GetBackingStore().Get("authoringLanguages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LocaleInfoable)
    }
    return nil
}
// GetDefaultDisplayLanguage gets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetDefaultDisplayLanguage()(LocaleInfoable) {
    val, err := m.GetBackingStore().Get("defaultDisplayLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LocaleInfoable)
    }
    return nil
}
// GetDefaultRegionalFormat gets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultRegionalFormat()(LocaleInfoable) {
    val, err := m.GetBackingStore().Get("defaultRegionalFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LocaleInfoable)
    }
    return nil
}
// GetDefaultSpeechInputLanguage gets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultSpeechInputLanguage()(LocaleInfoable) {
    val, err := m.GetBackingStore().Get("defaultSpeechInputLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LocaleInfoable)
    }
    return nil
}
// GetDefaultTranslationLanguage gets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultTranslationLanguage()(LocaleInfoable) {
    val, err := m.GetBackingStore().Get("defaultTranslationLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LocaleInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegionalAndLanguageSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authoringLanguages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocaleInfoable, len(val))
            for i, v := range val {
                res[i] = v.(LocaleInfoable)
            }
            m.SetAuthoringLanguages(res)
        }
        return nil
    }
    res["defaultDisplayLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDisplayLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["defaultRegionalFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRegionalFormat(val.(LocaleInfoable))
        }
        return nil
    }
    res["defaultSpeechInputLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSpeechInputLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["defaultTranslationLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultTranslationLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["regionalFormatOverrides"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRegionalFormatOverridesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionalFormatOverrides(val.(RegionalFormatOverridesable))
        }
        return nil
    }
    res["translationPreferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTranslationPreferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationPreferences(val.(TranslationPreferencesable))
        }
        return nil
    }
    return res
}
// GetRegionalFormatOverrides gets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) GetRegionalFormatOverrides()(RegionalFormatOverridesable) {
    val, err := m.GetBackingStore().Get("regionalFormatOverrides")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RegionalFormatOverridesable)
    }
    return nil
}
// GetTranslationPreferences gets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetTranslationPreferences()(TranslationPreferencesable) {
    val, err := m.GetBackingStore().Get("translationPreferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TranslationPreferencesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RegionalAndLanguageSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthoringLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthoringLanguages()))
        for i, v := range m.GetAuthoringLanguages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("authoringLanguages", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultDisplayLanguage sets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetDefaultDisplayLanguage(value LocaleInfoable)() {
    err := m.GetBackingStore().Set("defaultDisplayLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultRegionalFormat sets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultRegionalFormat(value LocaleInfoable)() {
    err := m.GetBackingStore().Set("defaultRegionalFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultSpeechInputLanguage sets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultSpeechInputLanguage(value LocaleInfoable)() {
    err := m.GetBackingStore().Set("defaultSpeechInputLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultTranslationLanguage sets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultTranslationLanguage(value LocaleInfoable)() {
    err := m.GetBackingStore().Set("defaultTranslationLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionalFormatOverrides sets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) SetRegionalFormatOverrides(value RegionalFormatOverridesable)() {
    err := m.GetBackingStore().Set("regionalFormatOverrides", value)
    if err != nil {
        panic(err)
    }
}
// SetTranslationPreferences sets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetTranslationPreferences(value TranslationPreferencesable)() {
    err := m.GetBackingStore().Set("translationPreferences", value)
    if err != nil {
        panic(err)
    }
}
// RegionalAndLanguageSettingsable 
type RegionalAndLanguageSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
