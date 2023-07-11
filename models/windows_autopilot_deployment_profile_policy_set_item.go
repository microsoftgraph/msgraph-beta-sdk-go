package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilePolicySetItem a class containing the properties used for windows autopilot deployment profile PolicySetItem.
type WindowsAutopilotDeploymentProfilePolicySetItem struct {
    PolicySetItem
}
// NewWindowsAutopilotDeploymentProfilePolicySetItem instantiates a new windowsAutopilotDeploymentProfilePolicySetItem and sets the default values.
func NewWindowsAutopilotDeploymentProfilePolicySetItem()(*WindowsAutopilotDeploymentProfilePolicySetItem) {
    m := &WindowsAutopilotDeploymentProfilePolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.windowsAutopilotDeploymentProfilePolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsAutopilotDeploymentProfilePolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilePolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilePolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfilePolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilePolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// WindowsAutopilotDeploymentProfilePolicySetItemable 
type WindowsAutopilotDeploymentProfilePolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
}
