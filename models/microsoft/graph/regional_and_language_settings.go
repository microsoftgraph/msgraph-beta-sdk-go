package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RegionalAndLanguageSettings struct {
    Entity
    // Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
    authoringLanguages []LocaleInfo;
    // The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
    defaultDisplayLanguage *LocaleInfo;
    // The locale that drives the default date, time, and calendar formatting.Returned by default.
    defaultRegionalFormat *LocaleInfo;
    // The language a user expected to use as input for text to speech scenarios.Returned by default.
    defaultSpeechInputLanguage *LocaleInfo;
    // The language a user expects to have documents, emails, and messages translated into.Returned by default.
    defaultTranslationLanguage *LocaleInfo;
    // Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
    regionalFormatOverrides *RegionalFormatOverrides;
    // The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
    translationPreferences *TranslationPreferences;
}
// Instantiates a new regionalAndLanguageSettings and sets the default values.
func NewRegionalAndLanguageSettings()(*RegionalAndLanguageSettings) {
    m := &RegionalAndLanguageSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetAuthoringLanguages()([]LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.authoringLanguages
    }
}
// Gets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetDefaultDisplayLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultDisplayLanguage
    }
}
// Gets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultRegionalFormat()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultRegionalFormat
    }
}
// Gets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultSpeechInputLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultSpeechInputLanguage
    }
}
// Gets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultTranslationLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultTranslationLanguage
    }
}
// Gets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) GetRegionalFormatOverrides()(*RegionalFormatOverrides) {
    if m == nil {
        return nil
    } else {
        return m.regionalFormatOverrides
    }
}
// Gets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetTranslationPreferences()(*TranslationPreferences) {
    if m == nil {
        return nil
    } else {
        return m.translationPreferences
    }
}
// The deserialization information for the current model
func (m *RegionalAndLanguageSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authoringLanguages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocaleInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*LocaleInfo))
            }
            m.SetAuthoringLanguages(res)
        }
        return nil
    }
    res["defaultDisplayLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDisplayLanguage(val.(*LocaleInfo))
        }
        return nil
    }
    res["defaultRegionalFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRegionalFormat(val.(*LocaleInfo))
        }
        return nil
    }
    res["defaultSpeechInputLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSpeechInputLanguage(val.(*LocaleInfo))
        }
        return nil
    }
    res["defaultTranslationLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultTranslationLanguage(val.(*LocaleInfo))
        }
        return nil
    }
    res["regionalFormatOverrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRegionalFormatOverrides() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionalFormatOverrides(val.(*RegionalFormatOverrides))
        }
        return nil
    }
    res["translationPreferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTranslationPreferences() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationPreferences(val.(*TranslationPreferences))
        }
        return nil
    }
    return res
}
func (m *RegionalAndLanguageSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RegionalAndLanguageSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthoringLanguages()))
        for i, v := range m.GetAuthoringLanguages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the authoringLanguages property.
func (m *RegionalAndLanguageSettings) SetAuthoringLanguages(value []LocaleInfo)() {
    m.authoringLanguages = value
}
// Sets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the defaultDisplayLanguage property.
func (m *RegionalAndLanguageSettings) SetDefaultDisplayLanguage(value *LocaleInfo)() {
    m.defaultDisplayLanguage = value
}
// Sets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
// Parameters:
//  - value : Value to set for the defaultRegionalFormat property.
func (m *RegionalAndLanguageSettings) SetDefaultRegionalFormat(value *LocaleInfo)() {
    m.defaultRegionalFormat = value
}
// Sets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
// Parameters:
//  - value : Value to set for the defaultSpeechInputLanguage property.
func (m *RegionalAndLanguageSettings) SetDefaultSpeechInputLanguage(value *LocaleInfo)() {
    m.defaultSpeechInputLanguage = value
}
// Sets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
// Parameters:
//  - value : Value to set for the defaultTranslationLanguage property.
func (m *RegionalAndLanguageSettings) SetDefaultTranslationLanguage(value *LocaleInfo)() {
    m.defaultTranslationLanguage = value
}
// Sets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
// Parameters:
//  - value : Value to set for the regionalFormatOverrides property.
func (m *RegionalAndLanguageSettings) SetRegionalFormatOverrides(value *RegionalFormatOverrides)() {
    m.regionalFormatOverrides = value
}
// Sets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the translationPreferences property.
func (m *RegionalAndLanguageSettings) SetTranslationPreferences(value *TranslationPreferences)() {
    m.translationPreferences = value
}
