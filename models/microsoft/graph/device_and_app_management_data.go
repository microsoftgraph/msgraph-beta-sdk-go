package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceAndAppManagementData struct {
    additionalData map[string]interface{};
    content []byte;
}
func NewDeviceAndAppManagementData()(*DeviceAndAppManagementData) {
    m := &DeviceAndAppManagementData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceAndAppManagementData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceAndAppManagementData) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *DeviceAndAppManagementData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementData) IsNil()(bool) {
    return m == nil
}
func (m *DeviceAndAppManagementData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
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
func (m *DeviceAndAppManagementData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceAndAppManagementData) SetContent(value []byte)() {
    m.content = value
}
