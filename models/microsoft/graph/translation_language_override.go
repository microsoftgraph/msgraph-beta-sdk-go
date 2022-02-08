package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TranslationLanguageOverride 
type TranslationLanguageOverride struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The language to apply the override.Returned by default. Not nullable.
    languageTag *string;
    // The translation override behavior for the language, if any.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior;
}
// NewTranslationLanguageOverride instantiates a new translationLanguageOverride and sets the default values.
func NewTranslationLanguageOverride()(*TranslationLanguageOverride) {
    m := &TranslationLanguageOverride{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationLanguageOverride) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetLanguageTag gets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetTranslationBehavior gets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslationLanguageOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["translationBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationBehavior(val.(*TranslationBehavior))
        }
        return nil
    }
    return res
}
func (m *TranslationLanguageOverride) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TranslationLanguageOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    if m.GetTranslationBehavior() != nil {
        cast := (*m.GetTranslationBehavior()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationLanguageOverride) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLanguageTag sets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetTranslationBehavior sets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) SetTranslationBehavior(value *TranslationBehavior)() {
    if m != nil {
        m.translationBehavior = value
    }
}
