package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MdmWindowsInformationProtectionPolicyPolicySetItem a class containing the properties used for mdm windows information protection policy PolicySetItem.
type MdmWindowsInformationProtectionPolicyPolicySetItem struct {
    PolicySetItem
}
// NewMdmWindowsInformationProtectionPolicyPolicySetItem instantiates a new MdmWindowsInformationProtectionPolicyPolicySetItem and sets the default values.
func NewMdmWindowsInformationProtectionPolicyPolicySetItem()(*MdmWindowsInformationProtectionPolicyPolicySetItem) {
    m := &MdmWindowsInformationProtectionPolicyPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.mdmWindowsInformationProtectionPolicyPolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMdmWindowsInformationProtectionPolicyPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMdmWindowsInformationProtectionPolicyPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMdmWindowsInformationProtectionPolicyPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MdmWindowsInformationProtectionPolicyPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MdmWindowsInformationProtectionPolicyPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type MdmWindowsInformationProtectionPolicyPolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
}
