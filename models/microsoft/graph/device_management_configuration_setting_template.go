package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSettingTemplate struct {
    Entity
    settingDefinitions []DeviceManagementConfigurationSettingDefinition;
    settingInstanceTemplate *DeviceManagementConfigurationSettingInstanceTemplate;
}
func NewDeviceManagementConfigurationSettingTemplate()(*DeviceManagementConfigurationSettingTemplate) {
    m := &DeviceManagementConfigurationSettingTemplate{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingInstanceTemplate()(*DeviceManagementConfigurationSettingInstanceTemplate) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplate
    }
}
func (m *DeviceManagementConfigurationSettingTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationSettingDefinition))
        }
        m.SetSettingDefinitions(res)
        return nil
    }
    res["settingInstanceTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstanceTemplate() })
        if err != nil {
            return err
        }
        m.SetSettingInstanceTemplate(val.(*DeviceManagementConfigurationSettingInstanceTemplate))
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingTemplate) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSettingTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settingInstanceTemplate", m.GetSettingInstanceTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinition)() {
    m.settingDefinitions = value
}
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingInstanceTemplate(value *DeviceManagementConfigurationSettingInstanceTemplate)() {
    m.settingInstanceTemplate = value
}
