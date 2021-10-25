package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDeviceCleanupSettings struct {
    additionalData map[string]interface{};
    deviceInactivityBeforeRetirementInDays *string;
}
func NewManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    m := &ManagedDeviceCleanupSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagedDeviceCleanupSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagedDeviceCleanupSettings) GetDeviceInactivityBeforeRetirementInDays()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDays
    }
}
func (m *ManagedDeviceCleanupSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceInactivityBeforeRetirementInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceInactivityBeforeRetirementInDays(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceCleanupSettings) IsNil()(bool) {
    return m == nil
}
func (m *ManagedDeviceCleanupSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceInactivityBeforeRetirementInDays", m.GetDeviceInactivityBeforeRetirementInDays())
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
func (m *ManagedDeviceCleanupSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagedDeviceCleanupSettings) SetDeviceInactivityBeforeRetirementInDays(value *string)() {
    m.deviceInactivityBeforeRetirementInDays = value
}
