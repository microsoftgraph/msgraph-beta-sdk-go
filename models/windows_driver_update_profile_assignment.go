package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateProfileAssignment this entity contains the properties used to assign a windows driver update profile to a group.
type WindowsDriverUpdateProfileAssignment struct {
    Entity
    // Base type for assignment targets.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewWindowsDriverUpdateProfileAssignment instantiates a new windowsDriverUpdateProfileAssignment and sets the default values.
func NewWindowsDriverUpdateProfileAssignment()(*WindowsDriverUpdateProfileAssignment) {
    m := &WindowsDriverUpdateProfileAssignment{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsDriverUpdateProfileAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsDriverUpdateProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateProfileAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfileAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetTarget gets the target property value. Base type for assignment targets.
func (m *WindowsDriverUpdateProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfileAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WindowsDriverUpdateProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
