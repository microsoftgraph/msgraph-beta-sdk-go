package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationPolicyAssignment struct {
    Entity
    // The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
    source *DeviceAndAppManagementAssignmentSource;
    // The identifier of the source of the assignment.
    sourceId *string;
    // The assignment target for the DeviceManagementConfigurationPolicy.
    target *DeviceAndAppManagementAssignmentTarget;
}
// Instantiates a new deviceManagementConfigurationPolicyAssignment and sets the default values.
func NewDeviceManagementConfigurationPolicyAssignment()(*DeviceManagementConfigurationPolicyAssignment) {
    m := &DeviceManagementConfigurationPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the source property value. The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// Gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// Gets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
func (m *DeviceManagementConfigurationPolicyAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceAndAppManagementAssignmentSource)
            m.SetSource(&cast)
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationPolicyAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationPolicyAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSource() != nil {
        cast := m.GetSource().String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the source property value. The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
// Parameters:
//  - value : Value to set for the source property.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    m.source = value
}
// Sets the sourceId property value. The identifier of the source of the assignment.
// Parameters:
//  - value : Value to set for the sourceId property.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSourceId(value *string)() {
    m.sourceId = value
}
// Sets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
// Parameters:
//  - value : Value to set for the target property.
func (m *DeviceManagementConfigurationPolicyAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
