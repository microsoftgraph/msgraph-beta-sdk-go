package wipemanagedappregistrationbydevicetag

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WipeManagedAppRegistrationByDeviceTagRequestBody struct {
    additionalData map[string]interface{};
    deviceTag *string;
}
func NewWipeManagedAppRegistrationByDeviceTagRequestBody()(*WipeManagedAppRegistrationByDeviceTagRequestBody) {
    m := &WipeManagedAppRegistrationByDeviceTagRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
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
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WipeManagedAppRegistrationByDeviceTagRequestBody) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
