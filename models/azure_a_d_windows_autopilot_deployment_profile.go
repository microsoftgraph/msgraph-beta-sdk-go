package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADWindowsAutopilotDeploymentProfile windows Autopilot Deployment Profile
type AzureADWindowsAutopilotDeploymentProfile struct {
    WindowsAutopilotDeploymentProfile
}
// NewAzureADWindowsAutopilotDeploymentProfile instantiates a new AzureADWindowsAutopilotDeploymentProfile and sets the default values.
func NewAzureADWindowsAutopilotDeploymentProfile()(*AzureADWindowsAutopilotDeploymentProfile) {
    m := &AzureADWindowsAutopilotDeploymentProfile{
        WindowsAutopilotDeploymentProfile: *NewWindowsAutopilotDeploymentProfile(),
    }
    odataTypeValue := "#microsoft.graph.azureADWindowsAutopilotDeploymentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureADWindowsAutopilotDeploymentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureADWindowsAutopilotDeploymentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADWindowsAutopilotDeploymentProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureADWindowsAutopilotDeploymentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsAutopilotDeploymentProfile.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureADWindowsAutopilotDeploymentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsAutopilotDeploymentProfile.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AzureADWindowsAutopilotDeploymentProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsAutopilotDeploymentProfileable
}
