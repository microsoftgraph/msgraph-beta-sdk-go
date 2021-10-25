package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingInstance struct {
    additionalData map[string]interface{};
    settingDefinitionId *string;
    settingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference;
}
func NewDeviceManagementConfigurationSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    m := &DeviceManagementConfigurationSettingInstance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementConfigurationSettingInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementConfigurationSettingInstance) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
func (m *DeviceManagementConfigurationSettingInstance) GetSettingInstanceTemplateReference()(*DeviceManagementConfigurationSettingInstanceTemplateReference) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateReference
    }
}
func (m *DeviceManagementConfigurationSettingInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["settingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDefinitionId(val)
        return nil
    }
    res["settingInstanceTemplateReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstanceTemplateReference() })
        if err != nil {
            return err
        }
        m.SetSettingInstanceTemplateReference(val.(*DeviceManagementConfigurationSettingInstanceTemplateReference))
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingInstance) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settingInstanceTemplateReference", m.GetSettingInstanceTemplateReference())
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
func (m *DeviceManagementConfigurationSettingInstance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementConfigurationSettingInstance) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
func (m *DeviceManagementConfigurationSettingInstance) SetSettingInstanceTemplateReference(value *DeviceManagementConfigurationSettingInstanceTemplateReference)() {
    m.settingInstanceTemplateReference = value
}
