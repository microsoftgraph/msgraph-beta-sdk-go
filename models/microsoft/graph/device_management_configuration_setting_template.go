package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationSettingTemplate struct {
    Entity
    // List of related Setting Definitions
    settingDefinitions []DeviceManagementConfigurationSettingDefinition;
    // Setting Instance Template
    settingInstanceTemplate *DeviceManagementConfigurationSettingInstanceTemplate;
}
// Instantiates a new deviceManagementConfigurationSettingTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingTemplate()(*DeviceManagementConfigurationSettingTemplate) {
    m := &DeviceManagementConfigurationSettingTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the settingDefinitions property value. List of related Setting Definitions
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// Gets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingInstanceTemplate()(*DeviceManagementConfigurationSettingInstanceTemplate) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplate
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the settingDefinitions property value. List of related Setting Definitions
// Parameters:
//  - value : Value to set for the settingDefinitions property.
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinition)() {
    m.settingDefinitions = value
}
// Sets the settingInstanceTemplate property value. Setting Instance Template
// Parameters:
//  - value : Value to set for the settingInstanceTemplate property.
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingInstanceTemplate(value *DeviceManagementConfigurationSettingInstanceTemplate)() {
    m.settingInstanceTemplate = value
}
