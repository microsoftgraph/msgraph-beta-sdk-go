package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationGroupAssignment device configuration group assignment.
type DeviceConfigurationGroupAssignment struct {
    Entity
    // The navigation link to the Device Configuration being targeted.
    deviceConfiguration DeviceConfigurationable
    // Indicates if this group is should be excluded. Defaults that the group should be included
    excludeGroup *bool
    // The Id of the AAD group we are targeting the device configuration to.
    targetGroupId *string
}
// NewDeviceConfigurationGroupAssignment instantiates a new deviceConfigurationGroupAssignment and sets the default values.
func NewDeviceConfigurationGroupAssignment()(*DeviceConfigurationGroupAssignment) {
    m := &DeviceConfigurationGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationGroupAssignment(), nil
}
// GetDeviceConfiguration gets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationGroupAssignment) GetDeviceConfiguration()(DeviceConfigurationable) {
    return m.deviceConfiguration
}
// GetExcludeGroup gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) GetExcludeGroup()(*bool) {
    return m.excludeGroup
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceConfigurationFromDiscriminatorValue , m.SetDeviceConfiguration)
    res["excludeGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetExcludeGroup)
    res["targetGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetGroupId)
    return res
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) GetTargetGroupId()(*string) {
    return m.targetGroupId
}
// Serialize serializes information the current object
func (m *DeviceConfigurationGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("excludeGroup", m.GetExcludeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceConfiguration sets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationGroupAssignment) SetDeviceConfiguration(value DeviceConfigurationable)() {
    m.deviceConfiguration = value
}
// SetExcludeGroup sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) SetExcludeGroup(value *bool)() {
    m.excludeGroup = value
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
