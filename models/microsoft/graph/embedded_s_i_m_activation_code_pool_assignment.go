package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EmbeddedSIMActivationCodePoolAssignment provides operations to manage the deviceManagement singleton.
type EmbeddedSIMActivationCodePoolAssignment struct {
    Entity
    // Base type for assignment targets.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewEmbeddedSIMActivationCodePoolAssignment instantiates a new embeddedSIMActivationCodePoolAssignment and sets the default values.
func NewEmbeddedSIMActivationCodePoolAssignment()(*EmbeddedSIMActivationCodePoolAssignment) {
    m := &EmbeddedSIMActivationCodePoolAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEmbeddedSIMActivationCodePoolAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCodePoolAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
// GetTarget gets the target property value. Base type for assignment targets.
func (m *EmbeddedSIMActivationCodePoolAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *EmbeddedSIMActivationCodePoolAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetTarget sets the target property value. Base type for assignment targets.
func (m *EmbeddedSIMActivationCodePoolAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
