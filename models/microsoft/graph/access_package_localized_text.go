package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageLocalizedText struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ISO code for the intended language. Required.
    languageCode *string;
    // The text in the specific language. Required.
    text *string;
}
// Instantiates a new accessPackageLocalizedText and sets the default values.
func NewAccessPackageLocalizedText()(*AccessPackageLocalizedText) {
    m := &AccessPackageLocalizedText{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageLocalizedText) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the languageCode property value. The ISO code for the intended language. Required.
func (m *AccessPackageLocalizedText) GetLanguageCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageCode
    }
}
// Gets the text property value. The text in the specific language. Required.
func (m *AccessPackageLocalizedText) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// The deserialization information for the current model
func (m *AccessPackageLocalizedText) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageCode(val)
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageLocalizedText) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageLocalizedText) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("languageCode", m.GetLanguageCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *AccessPackageLocalizedText) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the languageCode property value. The ISO code for the intended language. Required.
// Parameters:
//  - value : Value to set for the languageCode property.
func (m *AccessPackageLocalizedText) SetLanguageCode(value *string)() {
    m.languageCode = value
}
// Sets the text property value. The text in the specific language. Required.
// Parameters:
//  - value : Value to set for the text property.
func (m *AccessPackageLocalizedText) SetText(value *string)() {
    m.text = value
}
