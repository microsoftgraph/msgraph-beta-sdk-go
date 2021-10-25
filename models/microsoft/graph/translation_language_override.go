package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TranslationLanguageOverride struct {
    additionalData map[string]interface{};
    languageTag *string;
    translationBehavior *TranslationBehavior;
}
func NewTranslationLanguageOverride()(*TranslationLanguageOverride) {
    m := &TranslationLanguageOverride{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TranslationLanguageOverride) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TranslationLanguageOverride) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
func (m *TranslationLanguageOverride) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
func (m *TranslationLanguageOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguageTag(val)
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
    return res
}
func (m *TranslationLanguageOverride) IsNil()(bool) {
    return m == nil
}
func (m *TranslationLanguageOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TranslationLanguageOverride) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TranslationLanguageOverride) SetLanguageTag(value *string)() {
    m.languageTag = value
}
func (m *TranslationLanguageOverride) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
