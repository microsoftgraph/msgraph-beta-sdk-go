package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingOccurrence struct {
    additionalData map[string]interface{};
    maxDeviceOccurrence *int32;
    minDeviceOccurrence *int32;
}
func NewDeviceManagementConfigurationSettingOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    m := &DeviceManagementConfigurationSettingOccurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementConfigurationSettingOccurrence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementConfigurationSettingOccurrence) GetMaxDeviceOccurrence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxDeviceOccurrence
    }
}
func (m *DeviceManagementConfigurationSettingOccurrence) GetMinDeviceOccurrence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minDeviceOccurrence
    }
}
func (m *DeviceManagementConfigurationSettingOccurrence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maxDeviceOccurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxDeviceOccurrence(val)
        return nil
    }
    res["minDeviceOccurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinDeviceOccurrence(val)
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingOccurrence) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingOccurrence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maxDeviceOccurrence", m.GetMaxDeviceOccurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minDeviceOccurrence", m.GetMinDeviceOccurrence())
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
func (m *DeviceManagementConfigurationSettingOccurrence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementConfigurationSettingOccurrence) SetMaxDeviceOccurrence(value *int32)() {
    m.maxDeviceOccurrence = value
}
func (m *DeviceManagementConfigurationSettingOccurrence) SetMinDeviceOccurrence(value *int32)() {
    m.minDeviceOccurrence = value
}
