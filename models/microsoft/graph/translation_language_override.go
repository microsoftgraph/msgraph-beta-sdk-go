package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TranslationLanguageOverride struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The language to apply the override.Returned by default. Not nullable.
    languageTag *string;
    // The translation override behavior for the language, if any.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior;
}
// Instantiates a new translationLanguageOverride and sets the default values.
func NewTranslationLanguageOverride()(*TranslationLanguageOverride) {
    m := &TranslationLanguageOverride{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationLanguageOverride) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// Gets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
// The deserialization information for the current model
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
            cast := val.(TranslationBehavior)
            m.SetTranslationBehavior(&cast)
        }
        return nil
    }
    return res
}
func (m *TranslationLanguageOverride) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TranslationLanguageOverride) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the languageTag property.
func (m *TranslationLanguageOverride) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// Sets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the translationBehavior property.
func (m *TranslationLanguageOverride) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
