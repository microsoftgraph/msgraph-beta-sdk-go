package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyGroupAssignment device compliance policy group assignment.
type DeviceCompliancePolicyGroupAssignment struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeviceCompliancePolicyGroupAssignment instantiates a new deviceCompliancePolicyGroupAssignment and sets the default values.
func NewDeviceCompliancePolicyGroupAssignment()(*DeviceCompliancePolicyGroupAssignment) {
    m := &DeviceCompliancePolicyGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyGroupAssignment(), nil
}
// GetDeviceCompliancePolicy gets the deviceCompliancePolicy property value. The navigation link to the  device compliance polic targeted.
func (m *DeviceCompliancePolicyGroupAssignment) GetDeviceCompliancePolicy()(DeviceCompliancePolicyable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceCompliancePolicyable)
    }
    return nil
}
// GetExcludeGroup gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceCompliancePolicyGroupAssignment) GetExcludeGroup()(*bool) {
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
func (m *DeviceCompliancePolicyGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCompliancePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicy(val.(DeviceCompliancePolicyable))
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
// GetTargetGroupId gets the targetGroupId property value. The Id of the AAD group we are targeting the device compliance policy to.
func (m *DeviceCompliancePolicyGroupAssignment) GetTargetGroupId()(*string) {
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
func (m *DeviceCompliancePolicyGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicy", m.GetDeviceCompliancePolicy())
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
// SetDeviceCompliancePolicy sets the deviceCompliancePolicy property value. The navigation link to the  device compliance polic targeted.
func (m *DeviceCompliancePolicyGroupAssignment) SetDeviceCompliancePolicy(value DeviceCompliancePolicyable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeGroup sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceCompliancePolicyGroupAssignment) SetExcludeGroup(value *bool)() {
    err := m.GetBackingStore().Set("excludeGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the AAD group we are targeting the device compliance policy to.
func (m *DeviceCompliancePolicyGroupAssignment) SetTargetGroupId(value *string)() {
    err := m.GetBackingStore().Set("targetGroupId", value)
    if err != nil {
        panic(err)
    }
}
// DeviceCompliancePolicyGroupAssignmentable 
type DeviceCompliancePolicyGroupAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceCompliancePolicy()(DeviceCompliancePolicyable)
    GetExcludeGroup()(*bool)
    GetTargetGroupId()(*string)
    SetDeviceCompliancePolicy(value DeviceCompliancePolicyable)()
    SetExcludeGroup(value *bool)()
    SetTargetGroupId(value *string)()
}
