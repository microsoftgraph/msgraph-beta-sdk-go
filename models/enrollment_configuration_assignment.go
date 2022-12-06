package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentConfigurationAssignment enrollment Configuration Assignment
type EnrollmentConfigurationAssignment struct {
    Entity
    // Represents source of assignment.
    source *DeviceAndAppManagementAssignmentSource
    // Identifier for resource used for deployment to a group
    sourceId *string
    // Represents an assignment to managed devices in the tenant
    target DeviceAndAppManagementAssignmentTargetable
}
// NewEnrollmentConfigurationAssignment instantiates a new enrollmentConfigurationAssignment and sets the default values.
func NewEnrollmentConfigurationAssignment()(*EnrollmentConfigurationAssignment) {
    m := &EnrollmentConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentConfigurationAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentConfigurationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceAndAppManagementAssignmentSource , m.SetSource)
    res["sourceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSourceId)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetSource gets the source property value. Represents source of assignment.
func (m *EnrollmentConfigurationAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    return m.source
}
// GetSourceId gets the sourceId property value. Identifier for resource used for deployment to a group
func (m *EnrollmentConfigurationAssignment) GetSourceId()(*string) {
    return m.sourceId
}
// GetTarget gets the target property value. Represents an assignment to managed devices in the tenant
func (m *EnrollmentConfigurationAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *EnrollmentConfigurationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
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
// SetSource sets the source property value. Represents source of assignment.
func (m *EnrollmentConfigurationAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    m.source = value
}
// SetSourceId sets the sourceId property value. Identifier for resource used for deployment to a group
func (m *EnrollmentConfigurationAssignment) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTarget sets the target property value. Represents an assignment to managed devices in the tenant
func (m *EnrollmentConfigurationAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
