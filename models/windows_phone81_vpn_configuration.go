package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81VpnConfiguration 
type WindowsPhone81VpnConfiguration struct {
    Windows81VpnConfiguration
    // VPN Authentication Method.
    authenticationMethod *VpnAuthenticationMethod
    // Bypass VPN on company Wi-Fi.
    bypassVpnOnCompanyWifi *bool
    // Bypass VPN on home Wi-Fi.
    bypassVpnOnHomeWifi *bool
    // DNS suffix search list.
    dnsSuffixSearchList []string
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate WindowsPhone81CertificateProfileBaseable
    // Remember user credentials.
    rememberUserCredentials *bool
}
// NewWindowsPhone81VpnConfiguration instantiates a new WindowsPhone81VpnConfiguration and sets the default values.
func NewWindowsPhone81VpnConfiguration()(*WindowsPhone81VpnConfiguration) {
    m := &WindowsPhone81VpnConfiguration{
        Windows81VpnConfiguration: *NewWindows81VpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81VpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhone81VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81VpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *WindowsPhone81VpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetBypassVpnOnCompanyWifi gets the bypassVpnOnCompanyWifi property value. Bypass VPN on company Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) GetBypassVpnOnCompanyWifi()(*bool) {
    return m.bypassVpnOnCompanyWifi
}
// GetBypassVpnOnHomeWifi gets the bypassVpnOnHomeWifi property value. Bypass VPN on home Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) GetBypassVpnOnHomeWifi()(*bool) {
    return m.bypassVpnOnHomeWifi
}
// GetDnsSuffixSearchList gets the dnsSuffixSearchList property value. DNS suffix search list.
func (m *WindowsPhone81VpnConfiguration) GetDnsSuffixSearchList()([]string) {
    return m.dnsSuffixSearchList
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows81VpnConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVpnAuthenticationMethod , m.SetAuthenticationMethod)
    res["bypassVpnOnCompanyWifi"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBypassVpnOnCompanyWifi)
    res["bypassVpnOnHomeWifi"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBypassVpnOnHomeWifi)
    res["dnsSuffixSearchList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDnsSuffixSearchList)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsPhone81CertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["rememberUserCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRememberUserCredentials)
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *WindowsPhone81VpnConfiguration) GetIdentityCertificate()(WindowsPhone81CertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRememberUserCredentials gets the rememberUserCredentials property value. Remember user credentials.
func (m *WindowsPhone81VpnConfiguration) GetRememberUserCredentials()(*bool) {
    return m.rememberUserCredentials
}
// Serialize serializes information the current object
func (m *WindowsPhone81VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows81VpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bypassVpnOnCompanyWifi", m.GetBypassVpnOnCompanyWifi())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bypassVpnOnHomeWifi", m.GetBypassVpnOnHomeWifi())
        if err != nil {
            return err
        }
    }
    if m.GetDnsSuffixSearchList() != nil {
        err = writer.WriteCollectionOfStringValues("dnsSuffixSearchList", m.GetDnsSuffixSearchList())
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
    {
        err = writer.WriteBoolValue("rememberUserCredentials", m.GetRememberUserCredentials())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *WindowsPhone81VpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetBypassVpnOnCompanyWifi sets the bypassVpnOnCompanyWifi property value. Bypass VPN on company Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) SetBypassVpnOnCompanyWifi(value *bool)() {
    m.bypassVpnOnCompanyWifi = value
}
// SetBypassVpnOnHomeWifi sets the bypassVpnOnHomeWifi property value. Bypass VPN on home Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) SetBypassVpnOnHomeWifi(value *bool)() {
    m.bypassVpnOnHomeWifi = value
}
// SetDnsSuffixSearchList sets the dnsSuffixSearchList property value. DNS suffix search list.
func (m *WindowsPhone81VpnConfiguration) SetDnsSuffixSearchList(value []string)() {
    m.dnsSuffixSearchList = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *WindowsPhone81VpnConfiguration) SetIdentityCertificate(value WindowsPhone81CertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRememberUserCredentials sets the rememberUserCredentials property value. Remember user credentials.
func (m *WindowsPhone81VpnConfiguration) SetRememberUserCredentials(value *bool)() {
    m.rememberUserCredentials = value
}
