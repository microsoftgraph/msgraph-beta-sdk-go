package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentAssignment provides operations to manage the deviceManagement singleton.
type DeviceManagementIntentAssignment struct {
    Entity
    // The assignment target
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewDeviceManagementIntentAssignment instantiates a new deviceManagementIntentAssignment and sets the default values.
func NewDeviceManagementIntentAssignment()(*DeviceManagementIntentAssignment) {
    m := &DeviceManagementIntentAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementIntentAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
// GetTarget gets the target property value. The assignment target
func (m *DeviceManagementIntentAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *DeviceManagementIntentAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The assignment target
func (m *DeviceManagementIntentAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
