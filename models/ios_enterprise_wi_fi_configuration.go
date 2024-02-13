package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEnterpriseWiFiConfiguration by providing the configurations in this profile you can instruct the iOS device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user.
type IosEnterpriseWiFiConfiguration struct {
    IosWiFiConfiguration
}
// NewIosEnterpriseWiFiConfiguration instantiates a new IosEnterpriseWiFiConfiguration and sets the default values.
func NewIosEnterpriseWiFiConfiguration()(*IosEnterpriseWiFiConfiguration) {
    m := &IosEnterpriseWiFiConfiguration{
        IosWiFiConfiguration: *NewIosWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosEnterpriseWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
// returns a *WiFiAuthenticationMethod when successful
func (m *IosEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiAuthenticationMethod)
    }
    return nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
// returns a DeviceManagementDerivedCredentialSettingsable when successful
func (m *IosEnterpriseWiFiConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetEapFastConfiguration gets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
// returns a *EapFastConfiguration when successful
func (m *IosEnterpriseWiFiConfiguration) GetEapFastConfiguration()(*EapFastConfiguration) {
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
// returns a *EapType when successful
func (m *IosEnterpriseWiFiConfiguration) GetEapType()(*EapType) {
    val, err := m.GetBackingStore().Get("eapType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EapType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IosEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWiFiConfiguration.GetFieldDeserializers()
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
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
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
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEapTtls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEapTtls(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["outerIdentityPrivacyTemporaryValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuterIdentityPrivacyTemporaryValue(val)
        }
        return nil
    }
    res["passwordFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordFormatString(val)
        }
        return nil
    }
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosTrustedRootCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosTrustedRootCertificateable)
                }
            }
            m.SetRootCertificatesForServerValidation(res)
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
    res["usernameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameFormatString(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
// returns a IosCertificateProfileBaseable when successful
func (m *IosEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
// returns a *NonEapAuthenticationMethodForEapTtlsType when successful
func (m *IosEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEapTtls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
// returns a *string when successful
func (m *IosEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    val, err := m.GetBackingStore().Get("outerIdentityPrivacyTemporaryValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordFormatString gets the passwordFormatString property value. Password format string used to build the password to connect to wifi
// returns a *string when successful
func (m *IosEnterpriseWiFiConfiguration) GetPasswordFormatString()(*string) {
    val, err := m.GetBackingStore().Get("passwordFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
// returns a []IosTrustedRootCertificateable when successful
func (m *IosEnterpriseWiFiConfiguration) GetRootCertificatesForServerValidation()([]IosTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificatesForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosTrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
// returns a []string when successful
func (m *IosEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
    val, err := m.GetBackingStore().Get("trustedServerCertificateNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetUsernameFormatString gets the usernameFormatString property value. Username format string used to build the username to connect to wifi
// returns a *string when successful
func (m *IosEnterpriseWiFiConfiguration) GetUsernameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("usernameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWiFiConfiguration.Serialize(writer)
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
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
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
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetInnerAuthenticationProtocolForEapTtls() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEapTtls()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEapTtls", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outerIdentityPrivacyTemporaryValue", m.GetOuterIdentityPrivacyTemporaryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordFormatString", m.GetPasswordFormatString())
        if err != nil {
            return err
        }
    }
    if m.GetRootCertificatesForServerValidation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRootCertificatesForServerValidation()))
        for i, v := range m.GetRootCertificatesForServerValidation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rootCertificatesForServerValidation", cast)
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
    {
        err = writer.WriteStringValue("usernameFormatString", m.GetUsernameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *IosEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEnterpriseWiFiConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetEapFastConfiguration sets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *IosEnterpriseWiFiConfiguration) SetEapFastConfiguration(value *EapFastConfiguration)() {
    err := m.GetBackingStore().Set("eapFastConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *IosEnterpriseWiFiConfiguration) SetEapType(value *EapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *IosEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *IosEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEapTtls", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
func (m *IosEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordFormatString sets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) SetPasswordFormatString(value *string)() {
    err := m.GetBackingStore().Set("passwordFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
func (m *IosEnterpriseWiFiConfiguration) SetRootCertificatesForServerValidation(value []IosTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificatesForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *IosEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameFormatString sets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) SetUsernameFormatString(value *string)() {
    err := m.GetBackingStore().Set("usernameFormatString", value)
    if err != nil {
        panic(err)
    }
}
type IosEnterpriseWiFiConfigurationable interface {
    IosWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPasswordFormatString()(*string)
    GetRootCertificatesForServerValidation()([]IosTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUsernameFormatString()(*string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPasswordFormatString(value *string)()
    SetRootCertificatesForServerValidation(value []IosTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUsernameFormatString(value *string)()
}
