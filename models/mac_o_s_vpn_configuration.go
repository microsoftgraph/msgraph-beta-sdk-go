package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSVpnConfiguration by providing the configurations in this profile you can instruct the Mac device to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type MacOSVpnConfiguration struct {
    AppleVpnConfiguration
}
// NewMacOSVpnConfiguration instantiates a new MacOSVpnConfiguration and sets the default values.
func NewMacOSVpnConfiguration()(*MacOSVpnConfiguration) {
    m := &MacOSVpnConfiguration{
        AppleVpnConfiguration: *NewAppleVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSVpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSVpnConfiguration(), nil
}
// GetDeploymentChannel gets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
// returns a *AppleDeploymentChannel when successful
func (m *MacOSVpnConfiguration) GetDeploymentChannel()(*AppleDeploymentChannel) {
    val, err := m.GetBackingStore().Get("deploymentChannel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleDeploymentChannel)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleVpnConfiguration.GetFieldDeserializers()
    res["deploymentChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleDeploymentChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentChannel(val.(*AppleDeploymentChannel))
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(MacOSCertificateProfileBaseable))
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
// returns a MacOSCertificateProfileBaseable when successful
func (m *MacOSVpnConfiguration) GetIdentityCertificate()(MacOSCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSCertificateProfileBaseable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentChannel() != nil {
        cast := (*m.GetDeploymentChannel()).String()
        err = writer.WriteStringValue("deploymentChannel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentChannel sets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
func (m *MacOSVpnConfiguration) SetDeploymentChannel(value *AppleDeploymentChannel)() {
    err := m.GetBackingStore().Set("deploymentChannel", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *MacOSVpnConfiguration) SetIdentityCertificate(value MacOSCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
type MacOSVpnConfigurationable interface {
    AppleVpnConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentChannel()(*AppleDeploymentChannel)
    GetIdentityCertificate()(MacOSCertificateProfileBaseable)
    SetDeploymentChannel(value *AppleDeploymentChannel)()
    SetIdentityCertificate(value MacOSCertificateProfileBaseable)()
}
