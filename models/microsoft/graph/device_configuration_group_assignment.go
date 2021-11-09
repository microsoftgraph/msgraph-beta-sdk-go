package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceConfigurationGroupAssignment struct {
    Entity
    // The navigation link to the Device Configuration being targeted.
    deviceConfiguration *DeviceConfiguration;
    // Indicates if this group is should be excluded. Defaults that the group should be included
    excludeGroup *bool;
    // The Id of the AAD group we are targeting the device configuration to.
    targetGroupId *string;
}
// Instantiates a new deviceConfigurationGroupAssignment and sets the default values.
func NewDeviceConfigurationGroupAssignment()(*DeviceConfigurationGroupAssignment) {
    m := &DeviceConfigurationGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationGroupAssignment) GetDeviceConfiguration()(*DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
// Gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) GetExcludeGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroup
    }
}
// Gets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// The deserialization information for the current model
func (m *DeviceConfigurationGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfiguration(val.(*DeviceConfiguration))
        }
        return nil
    }
    res["excludeGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeGroup(val)
        }
        return nil
    }
    res["targetGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceConfigurationGroupAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceConfigurationGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("excludeGroup", m.GetExcludeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
// Parameters:
//  - value : Value to set for the deviceConfiguration property.
func (m *DeviceConfigurationGroupAssignment) SetDeviceConfiguration(value *DeviceConfiguration)() {
    m.deviceConfiguration = value
}
// Sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
// Parameters:
//  - value : Value to set for the excludeGroup property.
func (m *DeviceConfigurationGroupAssignment) SetExcludeGroup(value *bool)() {
    m.excludeGroup = value
}
// Sets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
// Parameters:
//  - value : Value to set for the targetGroupId property.
func (m *DeviceConfigurationGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
