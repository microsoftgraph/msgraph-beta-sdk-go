package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementConfigurationSetting struct {
    Entity
    settingDefinitions []DeviceManagementConfigurationSettingDefinition;
    settingInstance *DeviceManagementConfigurationSettingInstance;
}
func NewDeviceManagementConfigurationSetting()(*DeviceManagementConfigurationSetting) {
    m := &DeviceManagementConfigurationSetting{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementConfigurationSetting) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
func (m *DeviceManagementConfigurationSetting) GetSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settingInstance
    }
}
func (m *DeviceManagementConfigurationSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["settingInstance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstance() })
        if err != nil {
            return err
        }
        m.SetSettingInstance(val.(*DeviceManagementConfigurationSettingInstance))
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSetting) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementConfigurationSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("settingInstance", m.GetSettingInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementConfigurationSetting) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinition)() {
    m.settingDefinitions = value
}
func (m *DeviceManagementConfigurationSetting) SetSettingInstance(value *DeviceManagementConfigurationSettingInstance)() {
    m.settingInstance = value
}
