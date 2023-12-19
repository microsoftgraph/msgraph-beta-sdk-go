package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerEnterpriseWiFiConfiguration by providing the configurations in this profile you can instruct the AOSP Device Owner device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user.
type AospDeviceOwnerEnterpriseWiFiConfiguration struct {
    AospDeviceOwnerWiFiConfiguration
}
// NewAospDeviceOwnerEnterpriseWiFiConfiguration instantiates a new aospDeviceOwnerEnterpriseWiFiConfiguration and sets the default values.
func NewAospDeviceOwnerEnterpriseWiFiConfiguration()(*AospDeviceOwnerEnterpriseWiFiConfiguration) {
    m := &AospDeviceOwnerEnterpriseWiFiConfiguration{
        AospDeviceOwnerWiFiConfiguration: *NewAospDeviceOwnerWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.aospDeviceOwnerEnterpriseWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAospDeviceOwnerEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAospDeviceOwnerEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAospDeviceOwnerEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod)
    }
    return nil
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetEapType()(*AndroidEapType) {
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
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AospDeviceOwnerWiFiConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod))
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
        val, err := n.GetObjectValue(CreateAospDeviceOwnerCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(AospDeviceOwnerCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEapTtls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEapTtls(val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls))
        }
        return nil
    }
    res["innerAuthenticationProtocolForPeap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForPeap(val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap))
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
    res["rootCertificateForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAospDeviceOwnerTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForServerValidation(val.(AospDeviceOwnerTrustedRootCertificateable))
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
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(AospDeviceOwnerCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AospDeviceOwnerCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEapTtls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls)
    }
    return nil
}
// GetInnerAuthenticationProtocolForPeap gets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. This collection can contain a maximum of 500 elements. Possible values are: none, microsoftChapVersionTwo.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForPeap()(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForPeap")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    val, err := m.GetBackingStore().Get("outerIdentityPrivacyTemporaryValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetRootCertificateForServerValidation()(AospDeviceOwnerTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AospDeviceOwnerTrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
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
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AospDeviceOwnerWiFiConfiguration.Serialize(writer)
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
// SetAuthenticationMethod sets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetEapType(value *AndroidEapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value AospDeviceOwnerCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEapTtls", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForPeap sets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. This collection can contain a maximum of 500 elements. Possible values are: none, microsoftChapVersionTwo.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForPeap(value *AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForPeap", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetRootCertificateForServerValidation(value AospDeviceOwnerTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AospDeviceOwnerEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// AospDeviceOwnerEnterpriseWiFiConfigurationable 
type AospDeviceOwnerEnterpriseWiFiConfigurationable interface {
    AospDeviceOwnerWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AospDeviceOwnerCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls)
    GetInnerAuthenticationProtocolForPeap()(*AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(AospDeviceOwnerTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *AospDeviceOwnerEnterpriseWiFiConfiguration_authenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AospDeviceOwnerCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForEapTtls)()
    SetInnerAuthenticationProtocolForPeap(value *AospDeviceOwnerEnterpriseWiFiConfiguration_innerAuthenticationProtocolForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value AospDeviceOwnerTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
