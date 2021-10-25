package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingInstanceTemplate struct {
    additionalData map[string]interface{};
    isRequired *bool;
    settingDefinitionId *string;
    settingInstanceTemplateId *string;
}
func NewDeviceManagementConfigurationSettingInstanceTemplate()(*DeviceManagementConfigurationSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationSettingInstanceTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingInstanceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateId
    }
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRequired(val)
        return nil
    }
    res["settingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDefinitionId(val)
        return nil
    }
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
func (m *DeviceManagementConfigurationSettingInstanceTemplate) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
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
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetIsRequired(value *bool)() {
    m.isRequired = value
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingInstanceTemplateId(value *string)() {
    m.settingInstanceTemplateId = value
}
