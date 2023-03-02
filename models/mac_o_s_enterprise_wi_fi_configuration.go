package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSEnterpriseWiFiConfiguration 
type MacOSEnterpriseWiFiConfiguration struct {
    MacOSWiFiConfiguration
}
// NewMacOSEnterpriseWiFiConfiguration instantiates a new MacOSEnterpriseWiFiConfiguration and sets the default values.
func NewMacOSEnterpriseWiFiConfiguration()(*MacOSEnterpriseWiFiConfiguration) {
    m := &MacOSEnterpriseWiFiConfiguration{
        MacOSWiFiConfiguration: *NewMacOSWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSEnterpriseWiFiConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *MacOSEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
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
func (m *MacOSEnterpriseWiFiConfiguration) GetEapFastConfiguration()(*EapFastConfiguration) {
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
func (m *MacOSEnterpriseWiFiConfiguration) GetEapType()(*EapType) {
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
func (m *MacOSEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSWiFiConfiguration.GetFieldDeserializers()
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
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSTrustedRootCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSTrustedRootCertificateable)
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
                res[i] = *(v.(*string))
            }
            m.SetTrustedServerCertificateNames(res)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *MacOSEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *MacOSEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEapTtls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
func (m *MacOSEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    val, err := m.GetBackingStore().Get("outerIdentityPrivacyTemporaryValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP.
func (m *MacOSEnterpriseWiFiConfiguration) GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSTrustedRootCertificateable)
    }
    return nil
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
func (m *MacOSEnterpriseWiFiConfiguration) GetRootCertificatesForServerValidation()([]MacOSTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificatesForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSTrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this Wi-Fi network.
func (m *MacOSEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
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
func (m *MacOSEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSWiFiConfiguration.Serialize(writer)
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
        err = writer.WriteObjectValue("rootCertificateForServerValidation", m.GetRootCertificateForServerValidation())
        if err != nil {
            return err
        }
    }
    if m.GetRootCertificatesForServerValidation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRootCertificatesForServerValidation()))
        for i, v := range m.GetRootCertificatesForServerValidation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *MacOSEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetEapFastConfiguration sets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *MacOSEnterpriseWiFiConfiguration) SetEapFastConfiguration(value *EapFastConfiguration)() {
    err := m.GetBackingStore().Set("eapFastConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *MacOSEnterpriseWiFiConfiguration) SetEapType(value *EapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *MacOSEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *MacOSEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEapTtls", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
func (m *MacOSEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP.
func (m *MacOSEnterpriseWiFiConfiguration) SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
func (m *MacOSEnterpriseWiFiConfiguration) SetRootCertificatesForServerValidation(value []MacOSTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificatesForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this Wi-Fi network.
func (m *MacOSEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
// MacOSEnterpriseWiFiConfigurationable 
type MacOSEnterpriseWiFiConfigurationable interface {
    MacOSWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable)
    GetRootCertificatesForServerValidation()([]MacOSTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)()
    SetRootCertificatesForServerValidation(value []MacOSTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
