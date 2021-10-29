package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EmbeddedSIMActivationCodePoolAssignment struct {
    Entity
    // Base type for assignment targets.
    target *DeviceAndAppManagementAssignmentTarget;
}
// Instantiates a new embeddedSIMActivationCodePoolAssignment and sets the default values.
func NewEmbeddedSIMActivationCodePoolAssignment()(*EmbeddedSIMActivationCodePoolAssignment) {
    m := &EmbeddedSIMActivationCodePoolAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the target property value. Base type for assignment targets.
func (m *EmbeddedSIMActivationCodePoolAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *EmbeddedSIMActivationCodePoolAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        return nil
    }
    return res
}
func (m *EmbeddedSIMActivationCodePoolAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EmbeddedSIMActivationCodePoolAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the target property value. Base type for assignment targets.
// Parameters:
//  - value : Value to set for the target property.
func (m *EmbeddedSIMActivationCodePoolAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
