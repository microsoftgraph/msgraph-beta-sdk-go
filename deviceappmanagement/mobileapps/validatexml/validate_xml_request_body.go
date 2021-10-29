package validatexml

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ValidateXmlRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    officeConfigurationXml []byte;
}
// Instantiates a new validateXmlRequestBody and sets the default values.
func NewValidateXmlRequestBody()(*ValidateXmlRequestBody) {
    m := &ValidateXmlRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateXmlRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the officeConfigurationXml property value. 
func (m *ValidateXmlRequestBody) GetOfficeConfigurationXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationXml
    }
}
// The deserialization information for the current model
func (m *ValidateXmlRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["officeConfigurationXml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetOfficeConfigurationXml(val)
        return nil
    }
    return res
}
func (m *ValidateXmlRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ValidateXmlRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("officeConfigurationXml", m.GetOfficeConfigurationXml())
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
func (m *ValidateXmlRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the officeConfigurationXml property value. 
// Parameters:
//  - value : Value to set for the officeConfigurationXml property.
func (m *ValidateXmlRequestBody) SetOfficeConfigurationXml(value []byte)() {
    m.officeConfigurationXml = value
}
