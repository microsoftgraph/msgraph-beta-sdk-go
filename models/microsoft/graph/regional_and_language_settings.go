package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RegionalAndLanguageSettings struct {
    Entity
    authoringLanguages []LocaleInfo;
    defaultDisplayLanguage *LocaleInfo;
    defaultRegionalFormat *LocaleInfo;
    defaultSpeechInputLanguage *LocaleInfo;
    defaultTranslationLanguage *LocaleInfo;
    regionalFormatOverrides *RegionalFormatOverrides;
    translationPreferences *TranslationPreferences;
}
func NewRegionalAndLanguageSettings()(*RegionalAndLanguageSettings) {
    m := &RegionalAndLanguageSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *RegionalAndLanguageSettings) GetAuthoringLanguages()([]LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.authoringLanguages
    }
}
func (m *RegionalAndLanguageSettings) GetDefaultDisplayLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultDisplayLanguage
    }
}
func (m *RegionalAndLanguageSettings) GetDefaultRegionalFormat()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultRegionalFormat
    }
}
func (m *RegionalAndLanguageSettings) GetDefaultSpeechInputLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultSpeechInputLanguage
    }
}
func (m *RegionalAndLanguageSettings) GetDefaultTranslationLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.defaultTranslationLanguage
    }
}
func (m *RegionalAndLanguageSettings) GetRegionalFormatOverrides()(*RegionalFormatOverrides) {
    if m == nil {
        return nil
    } else {
        return m.regionalFormatOverrides
    }
}
func (m *RegionalAndLanguageSettings) GetTranslationPreferences()(*TranslationPreferences) {
    if m == nil {
        return nil
    } else {
        return m.translationPreferences
    }
}
func (m *RegionalAndLanguageSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authoringLanguages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        res := make([]LocaleInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*LocaleInfo))
        }
        m.SetAuthoringLanguages(res)
        return nil
    }
    res["defaultDisplayLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetDefaultDisplayLanguage(val.(*LocaleInfo))
        return nil
    }
    res["defaultRegionalFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetDefaultRegionalFormat(val.(*LocaleInfo))
        return nil
    }
    res["defaultSpeechInputLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetDefaultSpeechInputLanguage(val.(*LocaleInfo))
        return nil
    }
    res["defaultTranslationLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetDefaultTranslationLanguage(val.(*LocaleInfo))
        return nil
    }
    res["regionalFormatOverrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRegionalFormatOverrides() })
        if err != nil {
            return err
        }
        m.SetRegionalFormatOverrides(val.(*RegionalFormatOverrides))
        return nil
    }
    res["translationPreferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTranslationPreferences() })
        if err != nil {
            return err
        }
        m.SetTranslationPreferences(val.(*TranslationPreferences))
        return nil
    }
    return res
}
func (m *RegionalAndLanguageSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *RegionalAndLanguageSettings) SetAuthoringLanguages(value []LocaleInfo)() {
    m.authoringLanguages = value
}
func (m *RegionalAndLanguageSettings) SetDefaultDisplayLanguage(value *LocaleInfo)() {
    m.defaultDisplayLanguage = value
}
func (m *RegionalAndLanguageSettings) SetDefaultRegionalFormat(value *LocaleInfo)() {
    m.defaultRegionalFormat = value
}
func (m *RegionalAndLanguageSettings) SetDefaultSpeechInputLanguage(value *LocaleInfo)() {
    m.defaultSpeechInputLanguage = value
}
func (m *RegionalAndLanguageSettings) SetDefaultTranslationLanguage(value *LocaleInfo)() {
    m.defaultTranslationLanguage = value
}
func (m *RegionalAndLanguageSettings) SetRegionalFormatOverrides(value *RegionalFormatOverrides)() {
    m.regionalFormatOverrides = value
}
func (m *RegionalAndLanguageSettings) SetTranslationPreferences(value *TranslationPreferences)() {
    m.translationPreferences = value
}
