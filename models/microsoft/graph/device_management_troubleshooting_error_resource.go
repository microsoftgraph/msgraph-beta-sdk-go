package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementTroubleshootingErrorResource struct {
    additionalData map[string]interface{};
    link *string;
    text *string;
}
func NewDeviceManagementTroubleshootingErrorResource()(*DeviceManagementTroubleshootingErrorResource) {
    m := &DeviceManagementTroubleshootingErrorResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementTroubleshootingErrorResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementTroubleshootingErrorResource) GetLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.link
    }
}
func (m *DeviceManagementTroubleshootingErrorResource) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *DeviceManagementTroubleshootingErrorResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["link"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLink(val)
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetText(val)
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingErrorResource) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementTroubleshootingErrorResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("link", m.GetLink())
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
func (m *DeviceManagementTroubleshootingErrorResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementTroubleshootingErrorResource) SetLink(value *string)() {
    m.link = value
}
func (m *DeviceManagementTroubleshootingErrorResource) SetText(value *string)() {
    m.text = value
}
