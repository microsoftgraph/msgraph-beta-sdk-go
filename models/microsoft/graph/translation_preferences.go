package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TranslationPreferences struct {
    additionalData map[string]interface{};
    languageOverrides []TranslationLanguageOverride;
    translationBehavior *TranslationBehavior;
    untranslatedLanguages []string;
}
func NewTranslationPreferences()(*TranslationPreferences) {
    m := &TranslationPreferences{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TranslationPreferences) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TranslationPreferences) GetLanguageOverrides()([]TranslationLanguageOverride) {
    if m == nil {
        return nil
    } else {
        return m.languageOverrides
    }
}
func (m *TranslationPreferences) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
func (m *TranslationPreferences) GetUntranslatedLanguages()([]string) {
    if m == nil {
        return nil
    } else {
        return m.untranslatedLanguages
    }
}
func (m *TranslationPreferences) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageOverrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTranslationLanguageOverride() })
        if err != nil {
            return err
        }
        res := make([]TranslationLanguageOverride, len(val))
        for i, v := range val {
            res[i] = *(v.(*TranslationLanguageOverride))
        }
        m.SetLanguageOverrides(res)
        return nil
    }
    res["translationBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        cast := val.(TranslationBehavior)
        m.SetTranslationBehavior(&cast)
        return nil
    }
    res["untranslatedLanguages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetUntranslatedLanguages(res)
        return nil
    }
    return res
}
func (m *TranslationPreferences) IsNil()(bool) {
    return m == nil
}
func (m *TranslationPreferences) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguageOverrides()))
        for i, v := range m.GetLanguageOverrides() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("languageOverrides", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTranslationBehavior() != nil {
        cast := m.GetTranslationBehavior().String()
        err := writer.WriteStringValue("translationBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("untranslatedLanguages", m.GetUntranslatedLanguages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TranslationPreferences) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TranslationPreferences) SetLanguageOverrides(value []TranslationLanguageOverride)() {
    m.languageOverrides = value
}
func (m *TranslationPreferences) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
func (m *TranslationPreferences) SetUntranslatedLanguages(value []string)() {
    m.untranslatedLanguages = value
}
