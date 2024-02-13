package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81VpnConfiguration by providing the configurations in this profile you can instruct the Windows Phone 8.1 to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type WindowsPhone81VpnConfiguration struct {
    Windows81VpnConfiguration
}
// NewWindowsPhone81VpnConfiguration instantiates a new WindowsPhone81VpnConfiguration and sets the default values.
func NewWindowsPhone81VpnConfiguration()(*WindowsPhone81VpnConfiguration) {
    m := &WindowsPhone81VpnConfiguration{
        Windows81VpnConfiguration: *NewWindows81VpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81VpnConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhone81VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsPhone81VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81VpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
// returns a *VpnAuthenticationMethod when successful
func (m *WindowsPhone81VpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnAuthenticationMethod)
    }
    return nil
}
// GetBypassVpnOnCompanyWifi gets the bypassVpnOnCompanyWifi property value. Bypass VPN on company Wi-Fi.
// returns a *bool when successful
func (m *WindowsPhone81VpnConfiguration) GetBypassVpnOnCompanyWifi()(*bool) {
    val, err := m.GetBackingStore().Get("bypassVpnOnCompanyWifi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBypassVpnOnHomeWifi gets the bypassVpnOnHomeWifi property value. Bypass VPN on home Wi-Fi.
// returns a *bool when successful
func (m *WindowsPhone81VpnConfiguration) GetBypassVpnOnHomeWifi()(*bool) {
    val, err := m.GetBackingStore().Get("bypassVpnOnHomeWifi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDnsSuffixSearchList gets the dnsSuffixSearchList property value. DNS suffix search list.
// returns a []string when successful
func (m *WindowsPhone81VpnConfiguration) GetDnsSuffixSearchList()([]string) {
    val, err := m.GetBackingStore().Get("dnsSuffixSearchList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsPhone81VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows81VpnConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*VpnAuthenticationMethod))
        }
        return nil
    }
    res["bypassVpnOnCompanyWifi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassVpnOnCompanyWifi(val)
        }
        return nil
    }
    res["bypassVpnOnHomeWifi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassVpnOnHomeWifi(val)
        }
        return nil
    }
    res["dnsSuffixSearchList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDnsSuffixSearchList(res)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsPhone81CertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(WindowsPhone81CertificateProfileBaseable))
        }
        return nil
    }
    res["rememberUserCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRememberUserCredentials(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
// returns a WindowsPhone81CertificateProfileBaseable when successful
func (m *WindowsPhone81VpnConfiguration) GetIdentityCertificate()(WindowsPhone81CertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsPhone81CertificateProfileBaseable)
    }
    return nil
}
// GetRememberUserCredentials gets the rememberUserCredentials property value. Remember user credentials.
// returns a *bool when successful
func (m *WindowsPhone81VpnConfiguration) GetRememberUserCredentials()(*bool) {
    val, err := m.GetBackingStore().Get("rememberUserCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetBypassVpnOnCompanyWifi sets the bypassVpnOnCompanyWifi property value. Bypass VPN on company Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) SetBypassVpnOnCompanyWifi(value *bool)() {
    err := m.GetBackingStore().Set("bypassVpnOnCompanyWifi", value)
    if err != nil {
        panic(err)
    }
}
// SetBypassVpnOnHomeWifi sets the bypassVpnOnHomeWifi property value. Bypass VPN on home Wi-Fi.
func (m *WindowsPhone81VpnConfiguration) SetBypassVpnOnHomeWifi(value *bool)() {
    err := m.GetBackingStore().Set("bypassVpnOnHomeWifi", value)
    if err != nil {
        panic(err)
    }
}
// SetDnsSuffixSearchList sets the dnsSuffixSearchList property value. DNS suffix search list.
func (m *WindowsPhone81VpnConfiguration) SetDnsSuffixSearchList(value []string)() {
    err := m.GetBackingStore().Set("dnsSuffixSearchList", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *WindowsPhone81VpnConfiguration) SetIdentityCertificate(value WindowsPhone81CertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetRememberUserCredentials sets the rememberUserCredentials property value. Remember user credentials.
func (m *WindowsPhone81VpnConfiguration) SetRememberUserCredentials(value *bool)() {
    err := m.GetBackingStore().Set("rememberUserCredentials", value)
    if err != nil {
        panic(err)
    }
}
type WindowsPhone81VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows81VpnConfigurationable
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetBypassVpnOnCompanyWifi()(*bool)
    GetBypassVpnOnHomeWifi()(*bool)
    GetDnsSuffixSearchList()([]string)
    GetIdentityCertificate()(WindowsPhone81CertificateProfileBaseable)
    GetRememberUserCredentials()(*bool)
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetBypassVpnOnCompanyWifi(value *bool)()
    SetBypassVpnOnHomeWifi(value *bool)()
    SetDnsSuffixSearchList(value []string)()
    SetIdentityCertificate(value WindowsPhone81CertificateProfileBaseable)()
    SetRememberUserCredentials(value *bool)()
}
