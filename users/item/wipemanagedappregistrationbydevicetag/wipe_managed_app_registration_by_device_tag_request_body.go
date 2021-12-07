package wipemanagedappregistrationbydevicetag

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WipeManagedAppRegistrationByDeviceTagRequestBody 
type WipeManagedAppRegistrationByDeviceTagRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceTag *string;
}
// NewWipeManagedAppRegistrationByDeviceTagRequestBody instantiates a new wipeManagedAppRegistrationByDeviceTagRequestBody and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBody()(*WipeManagedAppRegistrationByDeviceTagRequestBody) {
    m := &WipeManagedAppRegistrationByDeviceTagRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceTag gets the deviceTag property value. 
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    return res
}
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceTag", m.GetDeviceTag())
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
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceTag sets the deviceTag property value. 
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetDeviceTag(value *string)() {
    if m != nil {
        m.deviceTag = value
    }
}
