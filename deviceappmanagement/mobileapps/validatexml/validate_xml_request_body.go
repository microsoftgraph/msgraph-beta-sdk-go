package validatexml

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ValidateXmlRequestBody 
type ValidateXmlRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    officeConfigurationXml []byte;
}
// NewValidateXmlRequestBody instantiates a new validateXmlRequestBody and sets the default values.
func NewValidateXmlRequestBody()(*ValidateXmlRequestBody) {
    m := &ValidateXmlRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateXmlRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOfficeConfigurationXml gets the officeConfigurationXml property value. 
func (m *ValidateXmlRequestBody) GetOfficeConfigurationXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationXml
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateXmlRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["officeConfigurationXml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeConfigurationXml(val)
        }
        return nil
    }
    return res
}
func (m *ValidateXmlRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateXmlRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOfficeConfigurationXml sets the officeConfigurationXml property value. 
func (m *ValidateXmlRequestBody) SetOfficeConfigurationXml(value []byte)() {
    if m != nil {
        m.officeConfigurationXml = value
    }
}
