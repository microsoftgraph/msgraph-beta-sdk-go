package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RegionalAndLanguageSettings 
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
// NewRegionalAndLanguageSettings instantiates a new regionalAndLanguageSettings and sets the default values.
func NewRegionalAndLanguageSettings()(*RegionalAndLanguageSettings) {
    m := &RegionalAndLanguageSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetAuthoringLanguages gets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetAuthoringLanguages()([]LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.authoringLanguages
    }
}
// GetDefaultDisplayLanguage gets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetDefaultDisplayLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultDisplayLanguage
    }
}
// GetDefaultRegionalFormat gets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultRegionalFormat()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultRegionalFormat
    }
}
// GetDefaultSpeechInputLanguage gets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultSpeechInputLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultSpeechInputLanguage
    }
}
// GetDefaultTranslationLanguage gets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) GetDefaultTranslationLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultTranslationLanguage
    }
}
// GetRegionalFormatOverrides gets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) GetRegionalFormatOverrides()(*RegionalFormatOverrides) {
    if m == nil {
        return nil
    } else {
        return m.regionalFormatOverrides
    }
}
// GetTranslationPreferences gets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) GetTranslationPreferences()(*TranslationPreferences) {
    if m == nil {
        return nil
    } else {
        return m.translationPreferences
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *RegionalAndLanguageSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthoringLanguages() != nil {
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
// SetAuthoringLanguages sets the authoringLanguages property value. Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetAuthoringLanguages(value []LocaleInfo)() {
    if m != nil {
        m.authoringLanguages = value
    }
}
// SetDefaultDisplayLanguage sets the defaultDisplayLanguage property value. The  user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web applications.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetDefaultDisplayLanguage(value *LocaleInfo)() {
    if m != nil {
        m.defaultDisplayLanguage = value
    }
}
// SetDefaultRegionalFormat sets the defaultRegionalFormat property value. The locale that drives the default date, time, and calendar formatting.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultRegionalFormat(value *LocaleInfo)() {
    if m != nil {
        m.defaultRegionalFormat = value
    }
}
// SetDefaultSpeechInputLanguage sets the defaultSpeechInputLanguage property value. The language a user expected to use as input for text to speech scenarios.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultSpeechInputLanguage(value *LocaleInfo)() {
    if m != nil {
        m.defaultSpeechInputLanguage = value
    }
}
// SetDefaultTranslationLanguage sets the defaultTranslationLanguage property value. The language a user expects to have documents, emails, and messages translated into.Returned by default.
func (m *RegionalAndLanguageSettings) SetDefaultTranslationLanguage(value *LocaleInfo)() {
    if m != nil {
        m.defaultTranslationLanguage = value
    }
}
// SetRegionalFormatOverrides sets the regionalFormatOverrides property value. Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
func (m *RegionalAndLanguageSettings) SetRegionalFormatOverrides(value *RegionalFormatOverrides)() {
    if m != nil {
        m.regionalFormatOverrides = value
    }
}
// SetTranslationPreferences sets the translationPreferences property value. The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by default. Not nullable.
func (m *RegionalAndLanguageSettings) SetTranslationPreferences(value *TranslationPreferences)() {
    if m != nil {
        m.translationPreferences = value
    }
}
