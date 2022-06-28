package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActiveDirectoryWindowsAutopilotDeploymentProfile 
type ActiveDirectoryWindowsAutopilotDeploymentProfile struct {
    WindowsAutopilotDeploymentProfile
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Configuration to join Active Directory domain
    domainJoinConfiguration WindowsDomainJoinConfigurationable
    // The Autopilot Hybrid Azure AD join flow will continue even if it does not establish domain controller connectivity during OOBE.
    hybridAzureADJoinSkipConnectivityCheck *bool
}
// NewActiveDirectoryWindowsAutopilotDeploymentProfile instantiates a new ActiveDirectoryWindowsAutopilotDeploymentProfile and sets the default values.
func NewActiveDirectoryWindowsAutopilotDeploymentProfile()(*ActiveDirectoryWindowsAutopilotDeploymentProfile) {
    m := &ActiveDirectoryWindowsAutopilotDeploymentProfile{
        WindowsAutopilotDeploymentProfile: *NewWindowsAutopilotDeploymentProfile(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActiveDirectoryWindowsAutopilotDeploymentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActiveDirectoryWindowsAutopilotDeploymentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActiveDirectoryWindowsAutopilotDeploymentProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDomainJoinConfiguration gets the domainJoinConfiguration property value. Configuration to join Active Directory domain
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) GetDomainJoinConfiguration()(WindowsDomainJoinConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.domainJoinConfiguration
    }
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
    if m == nil {
        return nil
    } else {
        return m.hybridAzureADJoinSkipConnectivityCheck
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDomainJoinConfiguration sets the domainJoinConfiguration property value. Configuration to join Active Directory domain
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) SetDomainJoinConfiguration(value WindowsDomainJoinConfigurationable)() {
    if m != nil {
        m.domainJoinConfiguration = value
    }
}
// SetHybridAzureADJoinSkipConnectivityCheck sets the hybridAzureADJoinSkipConnectivityCheck property value. The Autopilot Hybrid Azure AD join flow will continue even if it does not establish domain controller connectivity during OOBE.
func (m *ActiveDirectoryWindowsAutopilotDeploymentProfile) SetHybridAzureADJoinSkipConnectivityCheck(value *bool)() {
    if m != nil {
        m.hybridAzureADJoinSkipConnectivityCheck = value
    }
}
