package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleScopeTagAutoAssignment 
type RoleScopeTagAutoAssignment struct {
    Entity
    // The auto-assignment target for the specific Role Scope Tag.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewRoleScopeTagAutoAssignment instantiates a new roleScopeTagAutoAssignment and sets the default values.
func NewRoleScopeTagAutoAssignment()(*RoleScopeTagAutoAssignment) {
    m := &RoleScopeTagAutoAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRoleScopeTagAutoAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTagAutoAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The auto-assignment target for the specific Role Scope Tag.
func (m *RoleScopeTagAutoAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *RoleScopeTagAutoAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The auto-assignment target for the specific Role Scope Tag.
func (m *RoleScopeTagAutoAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
