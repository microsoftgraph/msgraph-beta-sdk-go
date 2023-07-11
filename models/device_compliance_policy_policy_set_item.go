package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyPolicySetItem a class containing the properties used for device compliance policy PolicySetItem.
type DeviceCompliancePolicyPolicySetItem struct {
    PolicySetItem
    // The OdataType property
    OdataType *string
}
// NewDeviceCompliancePolicyPolicySetItem instantiates a new deviceCompliancePolicyPolicySetItem and sets the default values.
func NewDeviceCompliancePolicyPolicySetItem()(*DeviceCompliancePolicyPolicySetItem) {
    m := &DeviceCompliancePolicyPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.deviceCompliancePolicyPolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceCompliancePolicyPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// DeviceCompliancePolicyPolicySetItemable 
type DeviceCompliancePolicyPolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
}
