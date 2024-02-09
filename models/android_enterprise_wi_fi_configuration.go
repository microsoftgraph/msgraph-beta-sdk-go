package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEnterpriseWiFiConfiguration by providing the configurations in this profile you can instruct the Android device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user.
type AndroidEnterpriseWiFiConfiguration struct {
    AndroidWiFiConfiguration
}
// NewAndroidEnterpriseWiFiConfiguration instantiates a new AndroidEnterpriseWiFiConfiguration and sets the default values.
func NewAndroidEnterpriseWiFiConfiguration()(*AndroidEnterpriseWiFiConfiguration) {
    m := &AndroidEnterpriseWiFiConfiguration{
        AndroidWiFiConfiguration: *NewAndroidWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidEnterpriseWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
// returns a *WiFiAuthenticationMethod when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiFiAuthenticationMethod)
    }
    return nil
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
// returns a *AndroidEapType when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetEapType()(*AndroidEapType) {
    val, err := m.GetBackingStore().Get("eapType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidEapType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidWiFiConfiguration.GetFieldDeserializers()
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
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*AndroidEapType))
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(AndroidCertificateProfileBaseable))
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
    res["innerAuthenticationProtocolForPeap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForPeap)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForPeap(val.(*NonEapAuthenticationMethodForPeap))
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
    res["preSharedKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreSharedKey(val)
        }
        return nil
    }
    res["rootCertificateForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForServerValidation(val.(AndroidTrustedRootCertificateable))
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
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
// returns a AndroidCertificateProfileBaseable when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(AndroidCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
// returns a *NonEapAuthenticationMethodForEapTtlsType when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEapTtls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetInnerAuthenticationProtocolForPeap gets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
// returns a *NonEapAuthenticationMethodForPeap when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForPeap")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForPeap)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
// returns a *string when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
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
func (m *AndroidEnterpriseWiFiConfiguration) GetPasswordFormatString()(*string) {
    val, err := m.GetBackingStore().Get("passwordFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreSharedKey gets the preSharedKey property value. PreSharedKey used to build the password to connect to wifi
// returns a *string when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetPreSharedKey()(*string) {
    val, err := m.GetBackingStore().Get("preSharedKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
// returns a AndroidTrustedRootCertificateable when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetRootCertificateForServerValidation()(AndroidTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidTrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
// returns a []string when successful
func (m *AndroidEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
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
func (m *AndroidEnterpriseWiFiConfiguration) GetUsernameFormatString()(*string) {
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
func (m *AndroidEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidWiFiConfiguration.Serialize(writer)
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
    if m.GetInnerAuthenticationProtocolForPeap() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForPeap()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForPeap", &cast)
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
    {
        err = writer.WriteStringValue("preSharedKey", m.GetPreSharedKey())
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
    {
        err = writer.WriteStringValue("usernameFormatString", m.GetUsernameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AndroidEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AndroidEnterpriseWiFiConfiguration) SetEapType(value *AndroidEapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AndroidEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value AndroidCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEapTtls", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForPeap sets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForPeap", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordFormatString sets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetPasswordFormatString(value *string)() {
    err := m.GetBackingStore().Set("passwordFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetPreSharedKey sets the preSharedKey property value. PreSharedKey used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetPreSharedKey(value *string)() {
    err := m.GetBackingStore().Set("preSharedKey", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AndroidEnterpriseWiFiConfiguration) SetRootCertificateForServerValidation(value AndroidTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameFormatString sets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetUsernameFormatString(value *string)() {
    err := m.GetBackingStore().Set("usernameFormatString", value)
    if err != nil {
        panic(err)
    }
}
type AndroidEnterpriseWiFiConfigurationable interface {
    AndroidWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AndroidCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPasswordFormatString()(*string)
    GetPreSharedKey()(*string)
    GetRootCertificateForServerValidation()(AndroidTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUsernameFormatString()(*string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AndroidCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPasswordFormatString(value *string)()
    SetPreSharedKey(value *string)()
    SetRootCertificateForServerValidation(value AndroidTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUsernameFormatString(value *string)()
}
