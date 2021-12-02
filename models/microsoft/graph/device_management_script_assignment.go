package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptAssignment 
type DeviceManagementScriptAssignment struct {
    Entity
    // The Id of the Azure Active Directory group we are targeting the script to.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewDeviceManagementScriptAssignment instantiates a new deviceManagementScriptAssignment and sets the default values.
func NewDeviceManagementScriptAssignment()(*DeviceManagementScriptAssignment) {
    m := &DeviceManagementScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTarget gets the target property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
func (m *DeviceManagementScriptAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
