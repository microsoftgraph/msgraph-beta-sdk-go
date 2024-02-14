package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWiredNetworkConfiguration this entity provides descriptions of the declared methods, properties and relationships exposed by the Wired Network CSP.
type WindowsWiredNetworkConfiguration struct {
    DeviceConfiguration
}
// NewWindowsWiredNetworkConfiguration instantiates a new WindowsWiredNetworkConfiguration and sets the default values.
func NewWindowsWiredNetworkConfiguration()(*WindowsWiredNetworkConfiguration) {
    m := &WindowsWiredNetworkConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWiredNetworkConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsWiredNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsWiredNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWiredNetworkConfiguration(), nil
}
// GetAuthenticationBlockPeriodInMinutes gets the authenticationBlockPeriodInMinutes property value. Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed authentication attempt.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationBlockPeriodInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("authenticationBlockPeriodInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
// returns a *WiredNetworkAuthenticationMethod when successful
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationMethod()(*WiredNetworkAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiredNetworkAuthenticationMethod)
    }
    return nil
}
// GetAuthenticationPeriodInSeconds gets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("authenticationPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationRetryDelayPeriodInSeconds gets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationRetryDelayPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("authenticationRetryDelayPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationType gets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest, unknownFutureValue.
// returns a *WiredNetworkAuthenticationType when successful
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationType()(*WiredNetworkAuthenticationType) {
    val, err := m.GetBackingStore().Get("authenticationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiredNetworkAuthenticationType)
    }
    return nil
}
// GetCacheCredentials gets the cacheCredentials property value. When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they connect. When FALSE, do not cache credentials. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetCacheCredentials()(*bool) {
    val, err := m.GetBackingStore().Get("cacheCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableUserPromptForServerValidation gets the disableUserPromptForServerValidation property value. When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetDisableUserPromptForServerValidation()(*bool) {
    val, err := m.GetBackingStore().Get("disableUserPromptForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEapolStartPeriodInSeconds gets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetEapolStartPeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("eapolStartPeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
// returns a *EapType when successful
func (m *WindowsWiredNetworkConfiguration) GetEapType()(*EapType) {
    val, err := m.GetBackingStore().Get("eapType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EapType)
    }
    return nil
}
// GetEnforce8021X gets the enforce8021X property value. When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication. When FALSE, 802.1X is not required. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetEnforce8021X()(*bool) {
    val, err := m.GetBackingStore().Get("enforce8021X")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsWiredNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationBlockPeriodInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationBlockPeriodInMinutes(val)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiredNetworkAuthenticationMethod))
        }
        return nil
    }
    res["authenticationPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationRetryDelayPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRetryDelayPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*WiredNetworkAuthenticationType))
        }
        return nil
    }
    res["cacheCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheCredentials(val)
        }
        return nil
    }
    res["disableUserPromptForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserPromptForServerValidation(val)
        }
        return nil
    }
    res["eapolStartPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapolStartPeriodInSeconds(val)
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
    res["enforce8021X"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforce8021X(val)
        }
        return nil
    }
    res["forceFIPSCompliance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceFIPSCompliance(val)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEAPTTLS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEAPTTLS(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["maximumAuthenticationFailures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuthenticationFailures(val)
        }
        return nil
    }
    res["maximumEAPOLStartMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumEAPOLStartMessages(val)
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
    res["performServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformServerValidation(val)
        }
        return nil
    }
    res["requireCryptographicBinding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCryptographicBinding(val)
        }
        return nil
    }
    res["rootCertificateForClientValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForClientValidation(val.(Windows81TrustedRootCertificateable))
        }
        return nil
    }
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows81TrustedRootCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Windows81TrustedRootCertificateable)
                }
            }
            m.SetRootCertificatesForServerValidation(res)
        }
        return nil
    }
    res["secondaryAuthenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryAuthenticationMethod(val.(*WiredNetworkAuthenticationMethod))
        }
        return nil
    }
    res["secondaryIdentityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryIdentityCertificateForClientAuthentication(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["secondaryRootCertificateForClientValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryRootCertificateForClientValidation(val.(Windows81TrustedRootCertificateable))
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
// GetForceFIPSCompliance gets the forceFIPSCompliance property value. When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetForceFIPSCompliance()(*bool) {
    val, err := m.GetBackingStore().Get("forceFIPSCompliance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
// returns a WindowsCertificateProfileBaseable when successful
func (m *WindowsWiredNetworkConfiguration) GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsCertificateProfileBaseable)
    }
    return nil
}
// GetInnerAuthenticationProtocolForEAPTTLS gets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
// returns a *NonEapAuthenticationMethodForEapTtlsType when successful
func (m *WindowsWiredNetworkConfiguration) GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType) {
    val, err := m.GetBackingStore().Get("innerAuthenticationProtocolForEAPTTLS")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonEapAuthenticationMethodForEapTtlsType)
    }
    return nil
}
// GetMaximumAuthenticationFailures gets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetMaximumAuthenticationFailures()(*int32) {
    val, err := m.GetBackingStore().Get("maximumAuthenticationFailures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumEAPOLStartMessages gets the maximumEAPOLStartMessages property value. Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
// returns a *int32 when successful
func (m *WindowsWiredNetworkConfiguration) GetMaximumEAPOLStartMessages()(*int32) {
    val, err := m.GetBackingStore().Get("maximumEAPOLStartMessages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
// returns a *string when successful
func (m *WindowsWiredNetworkConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    val, err := m.GetBackingStore().Get("outerIdentityPrivacyTemporaryValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPerformServerValidation gets the performServerValidation property value. When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP. When FALSE, the certificate is not validated. Default value is TRUE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetPerformServerValidation()(*bool) {
    val, err := m.GetBackingStore().Get("performServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequireCryptographicBinding gets the requireCryptographicBinding property value. When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic binding. Default value is TRUE.
// returns a *bool when successful
func (m *WindowsWiredNetworkConfiguration) GetRequireCryptographicBinding()(*bool) {
    val, err := m.GetBackingStore().Get("requireCryptographicBinding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRootCertificateForClientValidation gets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
// returns a Windows81TrustedRootCertificateable when successful
func (m *WindowsWiredNetworkConfiguration) GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificateForClientValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
// returns a []Windows81TrustedRootCertificateable when successful
func (m *WindowsWiredNetworkConfiguration) GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificatesForServerValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetSecondaryAuthenticationMethod gets the secondaryAuthenticationMethod property value. Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
// returns a *WiredNetworkAuthenticationMethod when successful
func (m *WindowsWiredNetworkConfiguration) GetSecondaryAuthenticationMethod()(*WiredNetworkAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("secondaryAuthenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WiredNetworkAuthenticationMethod)
    }
    return nil
}
// GetSecondaryIdentityCertificateForClientAuthentication gets the secondaryIdentityCertificateForClientAuthentication property value. Specify secondary identity certificate for client authentication.
// returns a WindowsCertificateProfileBaseable when successful
func (m *WindowsWiredNetworkConfiguration) GetSecondaryIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("secondaryIdentityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsCertificateProfileBaseable)
    }
    return nil
}
// GetSecondaryRootCertificateForClientValidation gets the secondaryRootCertificateForClientValidation property value. Specify secondary root certificate for client validation.
// returns a Windows81TrustedRootCertificateable when successful
func (m *WindowsWiredNetworkConfiguration) GetSecondaryRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("secondaryRootCertificateForClientValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Specify trusted server certificate names.
// returns a []string when successful
func (m *WindowsWiredNetworkConfiguration) GetTrustedServerCertificateNames()([]string) {
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
func (m *WindowsWiredNetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("authenticationBlockPeriodInMinutes", m.GetAuthenticationBlockPeriodInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("authenticationPeriodInSeconds", m.GetAuthenticationPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("authenticationRetryDelayPeriodInSeconds", m.GetAuthenticationRetryDelayPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationType() != nil {
        cast := (*m.GetAuthenticationType()).String()
        err = writer.WriteStringValue("authenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cacheCredentials", m.GetCacheCredentials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUserPromptForServerValidation", m.GetDisableUserPromptForServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("eapolStartPeriodInSeconds", m.GetEapolStartPeriodInSeconds())
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
        err = writer.WriteBoolValue("enforce8021X", m.GetEnforce8021X())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("forceFIPSCompliance", m.GetForceFIPSCompliance())
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
    if m.GetInnerAuthenticationProtocolForEAPTTLS() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEAPTTLS()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEAPTTLS", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAuthenticationFailures", m.GetMaximumAuthenticationFailures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumEAPOLStartMessages", m.GetMaximumEAPOLStartMessages())
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
        err = writer.WriteBoolValue("performServerValidation", m.GetPerformServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireCryptographicBinding", m.GetRequireCryptographicBinding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForClientValidation", m.GetRootCertificateForClientValidation())
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
    if m.GetSecondaryAuthenticationMethod() != nil {
        cast := (*m.GetSecondaryAuthenticationMethod()).String()
        err = writer.WriteStringValue("secondaryAuthenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("secondaryIdentityCertificateForClientAuthentication", m.GetSecondaryIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("secondaryRootCertificateForClientValidation", m.GetSecondaryRootCertificateForClientValidation())
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
// SetAuthenticationBlockPeriodInMinutes sets the authenticationBlockPeriodInMinutes property value. Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed authentication attempt.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationBlockPeriodInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("authenticationBlockPeriodInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationMethod(value *WiredNetworkAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationPeriodInSeconds sets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("authenticationPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationRetryDelayPeriodInSeconds sets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationRetryDelayPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("authenticationRetryDelayPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationType sets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationType(value *WiredNetworkAuthenticationType)() {
    err := m.GetBackingStore().Set("authenticationType", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheCredentials sets the cacheCredentials property value. When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they connect. When FALSE, do not cache credentials. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetCacheCredentials(value *bool)() {
    err := m.GetBackingStore().Set("cacheCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableUserPromptForServerValidation sets the disableUserPromptForServerValidation property value. When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetDisableUserPromptForServerValidation(value *bool)() {
    err := m.GetBackingStore().Set("disableUserPromptForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetEapolStartPeriodInSeconds sets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetEapolStartPeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("eapolStartPeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWiredNetworkConfiguration) SetEapType(value *EapType)() {
    err := m.GetBackingStore().Set("eapType", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforce8021X sets the enforce8021X property value. When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication. When FALSE, 802.1X is not required. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetEnforce8021X(value *bool)() {
    err := m.GetBackingStore().Set("enforce8021X", value)
    if err != nil {
        panic(err)
    }
}
// SetForceFIPSCompliance sets the forceFIPSCompliance property value. When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetForceFIPSCompliance(value *bool)() {
    err := m.GetBackingStore().Set("forceFIPSCompliance", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetInnerAuthenticationProtocolForEAPTTLS sets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWiredNetworkConfiguration) SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)() {
    err := m.GetBackingStore().Set("innerAuthenticationProtocolForEAPTTLS", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumAuthenticationFailures sets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) SetMaximumAuthenticationFailures(value *int32)() {
    err := m.GetBackingStore().Set("maximumAuthenticationFailures", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumEAPOLStartMessages sets the maximumEAPOLStartMessages property value. Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) SetMaximumEAPOLStartMessages(value *int32)() {
    err := m.GetBackingStore().Set("maximumEAPOLStartMessages", value)
    if err != nil {
        panic(err)
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWiredNetworkConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    err := m.GetBackingStore().Set("outerIdentityPrivacyTemporaryValue", value)
    if err != nil {
        panic(err)
    }
}
// SetPerformServerValidation sets the performServerValidation property value. When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP. When FALSE, the certificate is not validated. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) SetPerformServerValidation(value *bool)() {
    err := m.GetBackingStore().Set("performServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireCryptographicBinding sets the requireCryptographicBinding property value. When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic binding. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) SetRequireCryptographicBinding(value *bool)() {
    err := m.GetBackingStore().Set("requireCryptographicBinding", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateForClientValidation sets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificateForClientValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWiredNetworkConfiguration) SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificatesForServerValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryAuthenticationMethod sets the secondaryAuthenticationMethod property value. Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryAuthenticationMethod(value *WiredNetworkAuthenticationMethod)() {
    err := m.GetBackingStore().Set("secondaryAuthenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryIdentityCertificateForClientAuthentication sets the secondaryIdentityCertificateForClientAuthentication property value. Specify secondary identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("secondaryIdentityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryRootCertificateForClientValidation sets the secondaryRootCertificateForClientValidation property value. Specify secondary root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("secondaryRootCertificateForClientValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWiredNetworkConfiguration) SetTrustedServerCertificateNames(value []string)() {
    err := m.GetBackingStore().Set("trustedServerCertificateNames", value)
    if err != nil {
        panic(err)
    }
}
type WindowsWiredNetworkConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationBlockPeriodInMinutes()(*int32)
    GetAuthenticationMethod()(*WiredNetworkAuthenticationMethod)
    GetAuthenticationPeriodInSeconds()(*int32)
    GetAuthenticationRetryDelayPeriodInSeconds()(*int32)
    GetAuthenticationType()(*WiredNetworkAuthenticationType)
    GetCacheCredentials()(*bool)
    GetDisableUserPromptForServerValidation()(*bool)
    GetEapolStartPeriodInSeconds()(*int32)
    GetEapType()(*EapType)
    GetEnforce8021X()(*bool)
    GetForceFIPSCompliance()(*bool)
    GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType)
    GetMaximumAuthenticationFailures()(*int32)
    GetMaximumEAPOLStartMessages()(*int32)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPerformServerValidation()(*bool)
    GetRequireCryptographicBinding()(*bool)
    GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable)
    GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable)
    GetSecondaryAuthenticationMethod()(*WiredNetworkAuthenticationMethod)
    GetSecondaryIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable)
    GetSecondaryRootCertificateForClientValidation()(Windows81TrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationBlockPeriodInMinutes(value *int32)()
    SetAuthenticationMethod(value *WiredNetworkAuthenticationMethod)()
    SetAuthenticationPeriodInSeconds(value *int32)()
    SetAuthenticationRetryDelayPeriodInSeconds(value *int32)()
    SetAuthenticationType(value *WiredNetworkAuthenticationType)()
    SetCacheCredentials(value *bool)()
    SetDisableUserPromptForServerValidation(value *bool)()
    SetEapolStartPeriodInSeconds(value *int32)()
    SetEapType(value *EapType)()
    SetEnforce8021X(value *bool)()
    SetForceFIPSCompliance(value *bool)()
    SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetMaximumAuthenticationFailures(value *int32)()
    SetMaximumEAPOLStartMessages(value *int32)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPerformServerValidation(value *bool)()
    SetRequireCryptographicBinding(value *bool)()
    SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)()
    SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)()
    SetSecondaryAuthenticationMethod(value *WiredNetworkAuthenticationMethod)()
    SetSecondaryIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)()
    SetSecondaryRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
