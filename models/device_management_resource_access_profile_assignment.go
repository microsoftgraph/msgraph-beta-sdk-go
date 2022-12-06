package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementResourceAccessProfileAssignment entity that describes tenant level settings for derived credentials
type DeviceManagementResourceAccessProfileAssignment struct {
    Entity
    // The administrator intent for the assignment of the profile.
    intent *DeviceManagementResourceAccessProfileIntent
    // The identifier of the source of the assignment.
    sourceId *string
    // Base type for assignment targets.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewDeviceManagementResourceAccessProfileAssignment instantiates a new deviceManagementResourceAccessProfileAssignment and sets the default values.
func NewDeviceManagementResourceAccessProfileAssignment()(*DeviceManagementResourceAccessProfileAssignment) {
    m := &DeviceManagementResourceAccessProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementResourceAccessProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementResourceAccessProfileAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementResourceAccessProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementResourceAccessProfileAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementResourceAccessProfileIntent , m.SetIntent)
    res["sourceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSourceId)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetIntent gets the intent property value. The administrator intent for the assignment of the profile.
func (m *DeviceManagementResourceAccessProfileAssignment) GetIntent()(*DeviceManagementResourceAccessProfileIntent) {
    return m.intent
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementResourceAccessProfileAssignment) GetSourceId()(*string) {
    return m.sourceId
}
// GetTarget gets the target property value. Base type for assignment targets.
func (m *DeviceManagementResourceAccessProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *DeviceManagementResourceAccessProfileAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
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
// SetIntent sets the intent property value. The administrator intent for the assignment of the profile.
func (m *DeviceManagementResourceAccessProfileAssignment) SetIntent(value *DeviceManagementResourceAccessProfileIntent)() {
    m.intent = value
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementResourceAccessProfileAssignment) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTarget sets the target property value. Base type for assignment targets.
func (m *DeviceManagementResourceAccessProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
