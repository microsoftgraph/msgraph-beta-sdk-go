package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSetting 
type DeviceManagementConfigurationSetting struct {
    Entity
    // List of related Setting Definitions. This property is read-only.
    settingDefinitions []DeviceManagementConfigurationSettingDefinition;
    // Setting instance within policy
    settingInstance *DeviceManagementConfigurationSettingInstance;
}
// NewDeviceManagementConfigurationSetting instantiates a new deviceManagementConfigurationSetting and sets the default values.
func NewDeviceManagementConfigurationSetting()(*DeviceManagementConfigurationSetting) {
    m := &DeviceManagementConfigurationSetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// GetSettingInstance gets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) GetSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settingInstance
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationSettingDefinition))
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settingInstance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstance(val.(*DeviceManagementConfigurationSettingInstance))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingDefinitions() != nil {
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
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinition)() {
    if m != nil {
        m.settingDefinitions = value
    }
}
// SetSettingInstance sets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) SetSettingInstance(value *DeviceManagementConfigurationSettingInstance)() {
    if m != nil {
        m.settingInstance = value
    }
}
