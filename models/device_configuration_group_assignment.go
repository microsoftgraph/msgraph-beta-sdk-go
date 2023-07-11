package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationGroupAssignment device configuration group assignment.
type DeviceConfigurationGroupAssignment struct {
    Entity
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
    val, err := m.GetBackingStore().Get("deviceConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceConfigurationable)
    }
    return nil
}
// GetExcludeGroup gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) GetExcludeGroup()(*bool) {
    val, err := m.GetBackingStore().Get("excludeGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfiguration(val.(DeviceConfigurationable))
        }
        return nil
    }
    res["excludeGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeGroup(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["targetGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationGroupAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) GetTargetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("targetGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("deviceConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeGroup sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) SetExcludeGroup(value *bool)() {
    err := m.GetBackingStore().Set("excludeGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceConfigurationGroupAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) SetTargetGroupId(value *string)() {
    err := m.GetBackingStore().Set("targetGroupId", value)
    if err != nil {
        panic(err)
    }
}
// DeviceConfigurationGroupAssignmentable 
type DeviceConfigurationGroupAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceConfiguration()(DeviceConfigurationable)
    GetExcludeGroup()(*bool)
    GetOdataType()(*string)
    GetTargetGroupId()(*string)
    SetDeviceConfiguration(value DeviceConfigurationable)()
    SetExcludeGroup(value *bool)()
    SetOdataType(value *string)()
    SetTargetGroupId(value *string)()
}
