package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementSettingCategory struct {
    Entity
    // The category name
    displayName *string;
    // The category contains top level required setting
    hasRequiredSetting *bool;
    // The setting definitions this category contains
    settingDefinitions []DeviceManagementSettingDefinition;
}
// Instantiates a new deviceManagementSettingCategory and sets the default values.
func NewDeviceManagementSettingCategory()(*DeviceManagementSettingCategory) {
    m := &DeviceManagementSettingCategory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The category name
func (m *DeviceManagementSettingCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the hasRequiredSetting property value. The category contains top level required setting
func (m *DeviceManagementSettingCategory) GetHasRequiredSetting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasRequiredSetting
    }
}
// Gets the settingDefinitions property value. The setting definitions this category contains
func (m *DeviceManagementSettingCategory) GetSettingDefinitions()([]DeviceManagementSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// The deserialization information for the current model
func (m *DeviceManagementSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hasRequiredSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRequiredSetting(val)
        }
        return nil
    }
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingDefinition))
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementSettingCategory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the displayName property value. The category name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementSettingCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the hasRequiredSetting property value. The category contains top level required setting
// Parameters:
//  - value : Value to set for the hasRequiredSetting property.
func (m *DeviceManagementSettingCategory) SetHasRequiredSetting(value *bool)() {
    m.hasRequiredSetting = value
}
// Sets the settingDefinitions property value. The setting definitions this category contains
// Parameters:
//  - value : Value to set for the settingDefinitions property.
func (m *DeviceManagementSettingCategory) SetSettingDefinitions(value []DeviceManagementSettingDefinition)() {
    m.settingDefinitions = value
}
