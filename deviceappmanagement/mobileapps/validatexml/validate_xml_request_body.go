package validatexml

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ValidateXmlRequestBody struct {
    additionalData map[string]interface{};
    officeConfigurationXml []byte;
}
func NewValidateXmlRequestBody()(*ValidateXmlRequestBody) {
    m := &ValidateXmlRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ValidateXmlRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ValidateXmlRequestBody) GetOfficeConfigurationXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationXml
    }
}
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
func (m *ValidateXmlRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ValidateXmlRequestBody) SetOfficeConfigurationXml(value []byte)() {
    m.officeConfigurationXml = value
}
