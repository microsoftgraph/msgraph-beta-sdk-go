package wipemanagedappregistrationbydevicetag

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WipeManagedAppRegistrationByDeviceTagRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceTag *string;
}
// Instantiates a new wipeManagedAppRegistrationByDeviceTagRequestBody and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBody()(*WipeManagedAppRegistrationByDeviceTagRequestBody) {
    m := &WipeManagedAppRegistrationByDeviceTagRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceTag property value. 
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// The deserialization information for the current model
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceTag(val)
        return nil
    }
    return res
}
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceTag property value. 
// Parameters:
//  - value : Value to set for the deviceTag property.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
