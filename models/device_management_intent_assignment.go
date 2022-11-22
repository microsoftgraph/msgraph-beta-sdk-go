package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentAssignment intent assignment entity
type DeviceManagementIntentAssignment struct {
    Entity
    // The assignment target
    target DeviceAndAppManagementAssignmentTargetable
}
// NewDeviceManagementIntentAssignment instantiates a new deviceManagementIntentAssignment and sets the default values.
func NewDeviceManagementIntentAssignment()(*DeviceManagementIntentAssignment) {
    m := &DeviceManagementIntentAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetTarget gets the target property value. The assignment target
func (m *DeviceManagementIntentAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    m.target = value
}
