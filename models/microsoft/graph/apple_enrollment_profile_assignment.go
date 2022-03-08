package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppleEnrollmentProfileAssignment provides operations to manage the deviceManagement singleton.
type AppleEnrollmentProfileAssignment struct {
    Entity
    // The assignment target for the Apple user initiated deployment profile.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewAppleEnrollmentProfileAssignment instantiates a new appleEnrollmentProfileAssignment and sets the default values.
func NewAppleEnrollmentProfileAssignment()(*AppleEnrollmentProfileAssignment) {
    m := &AppleEnrollmentProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppleEnrollmentProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleEnrollmentProfileAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
// GetTarget gets the target property value. The assignment target for the Apple user initiated deployment profile.
func (m *AppleEnrollmentProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *AppleEnrollmentProfileAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppleEnrollmentProfileAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The assignment target for the Apple user initiated deployment profile.
func (m *AppleEnrollmentProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
