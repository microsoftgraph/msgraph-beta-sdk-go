package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActiveDirectoryWindowsAutopilotDeploymentProfile windows Autopilot Deployment Profile
type ActiveDirectoryWindowsAutopilotDeploymentProfile struct {
    WindowsAutopilotDeploymentProfile
}
// NewActiveDirectoryWindowsAutopilotDeploymentProfile instantiates a new activeDirectoryWindowsAutopilotDeploymentProfile and sets the default values.
func NewActiveDirectoryWindowsAutopilotDeploymentProfile()(*ActiveDirectoryWindowsAutopilotDeploymentProfile) {
    m := &ActiveDirectoryWindowsAutopilotDeploymentProfile{
        WindowsAutopilotDeploymentProfile: *NewWindowsAutopilotDeploymentProfile(),
    }
    odataTypeValue := "#microsoft.graph.activeDirectoryWindowsAutopilotDeploymentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateActiveDirectoryWindowsAutopilotDeploymentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActiveDirectoryWindowsAutopilotDeploymentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActiveDirectoryWindowsAutopilotDeploymentProfile(), nil
}
// GetDomainJoinConfiguration gets the domainJoinConfiguration property value. Configuration to join Active Directory domain
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) GetDomainJoinConfiguration()(WindowsDomainJoinConfigurationable) {
    val, err := m.GetBackingStore().Get("domainJoinConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsDomainJoinConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsAutopilotDeploymentProfile.GetFieldDeserializers()
    res["domainJoinConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsDomainJoinConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainJoinConfiguration(val.(WindowsDomainJoinConfigurationable))
        }
        return nil
    }
    res["hybridAzureADJoinSkipConnectivityCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHybridAzureADJoinSkipConnectivityCheck(val)
        }
        return nil
    }
    return res
}
// GetHybridAzureADJoinSkipConnectivityCheck gets the hybridAzureADJoinSkipConnectivityCheck property value. The Autopilot Hybrid Azure AD join flow will continue even if it does not establish domain controller connectivity during OOBE.
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) GetHybridAzureADJoinSkipConnectivityCheck()(*bool) {
    val, err := m.GetBackingStore().Get("hybridAzureADJoinSkipConnectivityCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsAutopilotDeploymentProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("domainJoinConfiguration", m.GetDomainJoinConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hybridAzureADJoinSkipConnectivityCheck", m.GetHybridAzureADJoinSkipConnectivityCheck())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDomainJoinConfiguration sets the domainJoinConfiguration property value. Configuration to join Active Directory domain
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) SetDomainJoinConfiguration(value WindowsDomainJoinConfigurationable)() {
    err := m.GetBackingStore().Set("domainJoinConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetHybridAzureADJoinSkipConnectivityCheck sets the hybridAzureADJoinSkipConnectivityCheck property value. The Autopilot Hybrid Azure AD join flow will continue even if it does not establish domain controller connectivity during OOBE.
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) SetHybridAzureADJoinSkipConnectivityCheck(value *bool)() {
    err := m.GetBackingStore().Set("hybridAzureADJoinSkipConnectivityCheck", value)
    if err != nil {
        panic(err)
    }
}
// ActiveDirectoryWindowsAutopilotDeploymentProfileable 
type ActiveDirectoryWindowsAutopilotDeploymentProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsAutopilotDeploymentProfileable
    GetDomainJoinConfiguration()(WindowsDomainJoinConfigurationable)
    GetHybridAzureADJoinSkipConnectivityCheck()(*bool)
    SetDomainJoinConfiguration(value WindowsDomainJoinConfigurationable)()
    SetHybridAzureADJoinSkipConnectivityCheck(value *bool)()
}
