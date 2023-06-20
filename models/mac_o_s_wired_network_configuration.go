package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSWiredNetworkConfiguration 
type MacOSWiredNetworkConfiguration struct {
    DeviceConfiguration
}
// NewMacOSWiredNetworkConfiguration instantiates a new MacOSWiredNetworkConfiguration and sets the default values.
func NewMacOSWiredNetworkConfiguration()(*MacOSWiredNetworkConfiguration) {
    m := &MacOSWiredNetworkConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSWiredNetworkConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSWiredNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSWiredNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSWiredNetworkConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *MacOSWiredNetworkConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiAuthenticationMethod)
    }
    return nil
}
// GetEapFastConfiguration gets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *MacOSWiredNetworkConfiguration) GetEapFastConfiguration()(*EapFastConfiguration) {
    val, err := m.GetBackingStore().Get("eapFastConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EapFastConfiguration)
    }
    return nil
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *MacOSWiredNetworkConfiguration) GetEapType()(*EapType) {
    val, err := m.GetBackingStore().Get("eapType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EapType)
    }
    return nil
}
// GetEnableOuterIdentityPrivacy gets the enableOuterIdentityPrivacy property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this wired network using their real username is displayed as 'anonymous'.
func (m *MacOSWiredNetworkConfiguration) GetEnableOuterIdentityPrivacy()(*string) {
    val, err := m.GetBackingStore().Get("enableOuterIdentityPrivacy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSWiredNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiFiAuthenticationMethod))
        }
        return nil
    }
    res["eapFastConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEapFastConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapFastConfiguration(val.(*EapFastConfiguration))
        }
        return nil
    }
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*EapType))
        }
        return nil
    }
    res["enableOuterIdentityPrivacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOuterIdentityPrivacy(val)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(MacOSCertificateProfileBaseable))
        }
        return nil
    }
    res["networkInterface"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkInterface)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInterface(val.(*WiredNetworkInterface))
        }
        return nil
    }
    res["networkName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkName(val)
        }
        return nil
    }
    res["nonEapAuthenticationMethodForEapTtls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonEapAuthenticationMethodForEapTtls(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["rootCertificateForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForServerValidation(val.(MacOSTrustedRootCertificateable))
        }
        return nil
    }
    res["trustedServerCertificateNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTrustedServerCertificateNames(res)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *MacOSWiredNetworkConfiguration) GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSCertificateProfileBaseable)
    }
    return nil
}
// GetNetworkInterface gets the networkInterface property value. Apple network interface type.
func (m *MacOSWiredNetworkConfiguration) GetNetworkInterface()(*WiredNetworkInterface) {
    val, err := m.GetBackingStore().Get("networkInterface")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiredNetworkInterface)
    }
    return nil
}
// GetNetworkName gets the networkName property value. Network Name
func (m *MacOSWiredNetworkConfiguration) GetNetworkName()(*string) {
    val, err := m.GetBackingStore().Get("networkName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNonEapAuthenticationMethodForEapTtls gets the nonEapAuthenticationMethodForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *MacOSWiredNetworkConfiguration) GetNonEapAuthenticationMethodForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("nonEapAuthenticationMethodForEapTtls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP.
func (m *MacOSWiredNetworkConfiguration) GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSTrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this wired network.
func (m *MacOSWiredNetworkConfiguration) GetTrustedServerCertificateNames()([]string) {
    val, err := m.GetBackingStore().Get("trustedServerCertificateNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSWiredNetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
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
    if m.GetEapFastConfiguration() != nil {
        cast := (*m.GetEapFastConfiguration()).String()
        err = writer.WriteStringValue("eapFastConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEapType() != nil {
        cast := (*m.GetEapType()).String()
        err = writer.WriteStringValue("eapType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enableOuterIdentityPrivacy", m.GetEnableOuterIdentityPrivacy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkInterface() != nil {
        cast := (*m.GetNetworkInterface()).String()
        err = writer.WriteStringValue("networkInterface", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkName", m.GetNetworkName())
        if err != nil {
            return err
        }
    }
    if m.GetNonEapAuthenticationMethodForEapTtls() != nil {
        cast := (*m.GetNonEapAuthenticationMethodForEapTtls()).String()
        err = writer.WriteStringValue("nonEapAuthenticationMethodForEapTtls", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForServerValidation", m.GetRootCertificateForServerValidation())
        if err != nil {
            return err
        }
    }
    if m.GetTrustedServerCertificateNames() != nil {
        err = writer.WriteCollectionOfStringValues("trustedServerCertificateNames", m.GetTrustedServerCertificateNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *MacOSWiredNetworkConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetEapFastConfiguration sets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *MacOSWiredNetworkConfiguration) SetEapFastConfiguration(value *EapFastConfiguration)() {
    err := m.GetBackingStore().Set("eapFastConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *MacOSWiredNetworkConfiguration) SetEapType(value *EapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableOuterIdentityPrivacy sets the enableOuterIdentityPrivacy property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this wired network using their real username is displayed as 'anonymous'.
func (m *MacOSWiredNetworkConfiguration) SetEnableOuterIdentityPrivacy(value *string)() {
    err := m.GetBackingStore().Set("enableOuterIdentityPrivacy", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *MacOSWiredNetworkConfiguration) SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkInterface sets the networkInterface property value. Apple network interface type.
func (m *MacOSWiredNetworkConfiguration) SetNetworkInterface(value *WiredNetworkInterface)() {
    err := m.GetBackingStore().Set("networkInterface", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkName sets the networkName property value. Network Name
func (m *MacOSWiredNetworkConfiguration) SetNetworkName(value *string)() {
    err := m.GetBackingStore().Set("networkName", value)
    if err != nil {
        panic(err)
    }
}
// SetNonEapAuthenticationMethodForEapTtls sets the nonEapAuthenticationMethodForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *MacOSWiredNetworkConfiguration) SetNonEapAuthenticationMethodForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("nonEapAuthenticationMethodForEapTtls", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP.
func (m *MacOSWiredNetworkConfiguration) SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this wired network.
func (m *MacOSWiredNetworkConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// MacOSWiredNetworkConfigurationable 
type MacOSWiredNetworkConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetEnableOuterIdentityPrivacy()(*string)
    GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable)
    GetNetworkInterface()(*WiredNetworkInterface)
    GetNetworkName()(*string)
    GetNonEapAuthenticationMethodForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetEnableOuterIdentityPrivacy(value *string)()
    SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)()
    SetNetworkInterface(value *WiredNetworkInterface)()
    SetNetworkName(value *string)()
    SetNonEapAuthenticationMethodForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
