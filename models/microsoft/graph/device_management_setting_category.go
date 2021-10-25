package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettingCategory struct {
    Entity
    displayName *string;
    hasRequiredSetting *bool;
    settingDefinitions []DeviceManagementSettingDefinition;
}
func NewDeviceManagementSettingCategory()(*DeviceManagementSettingCategory) {
    m := &DeviceManagementSettingCategory{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementSettingCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementSettingCategory) GetHasRequiredSetting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasRequiredSetting
    }
}
func (m *DeviceManagementSettingCategory) GetSettingDefinitions()([]DeviceManagementSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
func (m *DeviceManagementSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["hasRequiredSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasRequiredSetting(val)
        return nil
    }
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingDefinition))
        }
        m.SetSettingDefinitions(res)
        return nil
    }
    return res
}
func (m *DeviceManagementSettingCategory) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementSettingCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasRequiredSetting", m.GetHasRequiredSetting())
        if err != nil {
            return err
        }
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
    return nil
}
func (m *DeviceManagementSettingCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementSettingCategory) SetHasRequiredSetting(value *bool)() {
    m.hasRequiredSetting = value
}
func (m *DeviceManagementSettingCategory) SetSettingDefinitions(value []DeviceManagementSettingDefinition)() {
    m.settingDefinitions = value
}
