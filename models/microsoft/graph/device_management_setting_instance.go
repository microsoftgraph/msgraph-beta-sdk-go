package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingInstance 
type DeviceManagementSettingInstance struct {
    Entity
    // The ID of the setting definition for this instance
    definitionId *string;
    // JSON representation of the value
    valueJson *string;
}
// NewDeviceManagementSettingInstance instantiates a new deviceManagementSettingInstance and sets the default values.
func NewDeviceManagementSettingInstance()(*DeviceManagementSettingInstance) {
    m := &DeviceManagementSettingInstance{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefinitionId gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingInstance) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
// GetValueJson gets the valueJson property value. JSON representation of the value
func (m *DeviceManagementSettingInstance) GetValueJson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueJson
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDefinitionId sets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingInstance) SetDefinitionId(value *string)() {
    if m != nil {
        m.definitionId = value
    }
}
// SetValueJson sets the valueJson property value. JSON representation of the value
func (m *DeviceManagementSettingInstance) SetValueJson(value *string)() {
    if m != nil {
        m.valueJson = value
    }
}
