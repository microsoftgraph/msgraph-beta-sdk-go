package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementSettingInstance struct {
    Entity
    // The ID of the setting definition for this instance
    definitionId *string;
    // JSON representation of the value
    valueJson *string;
}
// Instantiates a new deviceManagementSettingInstance and sets the default values.
func NewDeviceManagementSettingInstance()(*DeviceManagementSettingInstance) {
    m := &DeviceManagementSettingInstance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingInstance) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
// Gets the valueJson property value. JSON representation of the value
func (m *DeviceManagementSettingInstance) GetValueJson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueJson
    }
}
// The deserialization information for the current model
func (m *DeviceManagementSettingInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionId(val)
        }
        return nil
    }
    res["valueJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueJson(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementSettingInstance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementSettingInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("definitionId", m.GetDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueJson", m.GetValueJson())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the definitionId property value. The ID of the setting definition for this instance
// Parameters:
//  - value : Value to set for the definitionId property.
func (m *DeviceManagementSettingInstance) SetDefinitionId(value *string)() {
    m.definitionId = value
}
// Sets the valueJson property value. JSON representation of the value
// Parameters:
//  - value : Value to set for the valueJson property.
func (m *DeviceManagementSettingInstance) SetValueJson(value *string)() {
    m.valueJson = value
}
