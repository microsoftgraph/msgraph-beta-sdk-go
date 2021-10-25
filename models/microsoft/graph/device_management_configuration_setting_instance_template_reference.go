package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingInstanceTemplateReference struct {
    additionalData map[string]interface{};
    settingInstanceTemplateId *string;
}
func NewDeviceManagementConfigurationSettingInstanceTemplateReference()(*DeviceManagementConfigurationSettingInstanceTemplateReference) {
    m := &DeviceManagementConfigurationSettingInstanceTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetSettingInstanceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateId
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["settingInstanceTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingInstanceTemplateId(val)
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingInstanceTemplateId", m.GetSettingInstanceTemplateId())
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
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetSettingInstanceTemplateId(value *string)() {
    m.settingInstanceTemplateId = value
}
